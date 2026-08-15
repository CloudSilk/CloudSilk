import React, { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import {
  Alert, Button, Card, DatePicker, Form, InputNumber, message, Select, Space,
  Spin, Table, Tag, Tooltip, Typography,
} from 'antd';
import { ReloadOutlined } from '@ant-design/icons';
import dayjs from 'dayjs';
import umiRequest from 'umi-request';
import { getToken } from '@swiftease/atali-pkg';

const { Title, Text } = Typography;

/** 后端统一响应码：Code_Success = 20000 */
const CODE_SUCCESS = 20000;

interface ScheduleItem {
  id: string;
  productOrderNo: string;
  sequence: number;
  plannedStartTime: string;
  plannedEndTime: string;
  durationSeconds: number;
}

interface SchedulePlan {
  id: string;
  planNo: string;
  productionLineID: string;
  productionLineCode: string;
  startTime: string;
  endTime: string;
  orderCount: number;
  currentState: string;
  items?: ScheduleItem[];
}

interface ProductionLine {
  id: string;
  code: string;
  description: string;
}

const authHeaders = (): Record<string, string> => ({
  'Content-Type': 'application/json',
  authorization: 'Bearer ' + getToken(),
});

/** 统一响应处理：成功返回整个 resp，失败抛 Error(message) */
const api = async (url: string, options: Record<string, any> = {}): Promise<any> => {
  const resp = await umiRequest<Record<string, any>>(url, { headers: authHeaders(), ...options });
  if (resp && (resp.code === CODE_SUCCESS || (resp.code === undefined && resp.data !== undefined))) {
    return resp;
  }
  throw new Error(resp?.message || `请求失败（code=${resp?.code}）`);
};

const fmt = 'YYYY-MM-DD HH:mm:ss';
const DRAG_THRESHOLD_PX = 4; // 低于该位移视为点击，不触发调整

/** 单条轨道：条形 + 拖拽交互（阈值、ESC 取消、钳制、实时预览） */
const GanttRow: React.FC<{
  item: ScheduleItem;
  color: string;
  rangeMin: number;
  rangeSpan: number;
  clampMinMs: number;
  adjustable: boolean;
  onAdjust: (item: ScheduleItem, newStartMs: number) => void;
}> = ({ item, color, rangeMin, rangeSpan, clampMinMs, adjustable, onAdjust }) => {
  const trackRef = useRef<HTMLDivElement | null>(null);
  const [drag, setDrag] = useState<{ x0: number; dx: number } | null>(null);

  const start = new Date(item.plannedStartTime).getTime();
  const end = new Date(item.plannedEndTime).getTime();
  const leftPct = ((start - rangeMin) / rangeSpan) * 100;
  const widthPct = Math.max(((end - start) / rangeSpan) * 100, 0.6);
  const offsetPx = drag?.dx ?? 0;
  const previewMs = start + (offsetPx / (trackRef.current?.clientWidth || 1)) * rangeSpan;
  const clamped = Math.max(previewMs, clampMinMs);

  useEffect(() => {
    if (!drag) return;
    const onMove = (e: MouseEvent) => setDrag((d) => (d ? { ...d, dx: e.clientX - d.x0 } : d));
    const onKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape') setDrag(null);
    };
    const onUp = () => {
      setDrag((d) => {
        if (d && Math.abs(d.dx) >= DRAG_THRESHOLD_PX) {
          const track = trackRef.current;
          if (track && track.clientWidth > 0) {
            const ms = (d.dx / track.clientWidth) * rangeSpan;
            onAdjust(item, Math.max(start + ms, clampMinMs));
          }
        }
        return null;
      });
    };
    window.addEventListener('mousemove', onMove);
    window.addEventListener('mouseup', onUp);
    window.addEventListener('keydown', onKey);
    return () => {
      window.removeEventListener('mousemove', onMove);
      window.removeEventListener('mouseup', onUp);
      window.removeEventListener('keydown', onKey);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [drag !== null]);

  const dragging = drag !== null && Math.abs(drag.dx) >= DRAG_THRESHOLD_PX;
  const tip = dragging
    ? `${dayjs(clamped).format('MM-DD HH:mm')} → ${dayjs(clamped + (end - start)).format('MM-DD HH:mm')}（松手保存，ESC 取消）`
    : `${item.plannedStartTime} ~ ${item.plannedEndTime}（${Math.round(item.durationSeconds / 60)} 分钟）${adjustable ? '，可拖拽调整' : ''}`;

  return (
    <div style={{ display: 'flex', alignItems: 'center', marginBottom: 4, gap: 8 }}>
      <Text style={{ width: 132, fontSize: 12, flexShrink: 0 }} ellipsis>
        {`${item.sequence}. ${item.productOrderNo}`}
      </Text>
      <div ref={trackRef} style={{ flex: 1, position: 'relative', height: 24, background: '#f0f0f0', borderRadius: 4, overflow: 'hidden' }}>
        <Tooltip title={tip} open={dragging || undefined}>
          <div
            onMouseDown={(e) => {
              if (!adjustable || e.button !== 0) return;
              setDrag({ x0: e.clientX, dx: 0 });
            }}
            style={{
              position: 'absolute',
              left: `calc(${leftPct}% + ${offsetPx}px)`,
              width: `${widthPct}%`,
              top: 4, bottom: 4,
              background: color, borderRadius: 3, minWidth: 6,
              cursor: adjustable ? 'ew-resize' : 'pointer',
              opacity: dragging ? 0.7 : 1,
              boxShadow: dragging ? '0 0 0 2px rgba(22,119,255,.4)' : undefined,
              transition: dragging ? 'none' : 'opacity .15s',
            }}
          />
        </Tooltip>
      </div>
      <Text style={{ width: 96, fontSize: 12, flexShrink: 0 }} type={dragging ? 'warning' : 'secondary'}>
        {dragging ? dayjs(clamped).format('MM-DD HH:mm') : dayjs(item.plannedStartTime).format('MM-DD HH:mm')}
      </Text>
    </div>
  );
};

/** 甘特图：时间刻度、现在线、条形拖拽 */
const Gantt: React.FC<{
  items: ScheduleItem[];
  planStartMs: number;
  adjustable: boolean;
  onAdjust: (item: ScheduleItem, newStartMs: number) => void;
}> = ({ items, planStartMs, adjustable, onAdjust }) => {
  const range = useMemo(() => {
    if (!items?.length) return null;
    const starts = items.map((i) => new Date(i.plannedStartTime).getTime()).filter((t) => !isNaN(t));
    const ends = items.map((i) => new Date(i.plannedEndTime).getTime()).filter((t) => !isNaN(t));
    if (!starts.length || !ends.length) return null;
    const min = Math.min(...starts);
    const max = Math.max(...ends);
    return { min, span: Math.max(max - min, 1) };
  }, [items]);

  if (!range) return <Alert type="info" message="该计划无明细" showIcon />;

  const colors = ['#1677ff', '#52c41a', '#fa8c16', '#722ed1', '#13c2c2', '#eb2f96'];
  const ticks = [0, 0.25, 0.5, 0.75, 1].map((p) => range.min + range.span * p);
  const now = Date.now();
  const nowPct = ((now - range.min) / range.span) * 100;
  const showNow = nowPct >= 0 && nowPct <= 100;

  return (
    <div>
      {/* 时间轴刻度 */}
      <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 2 }}>
        <div style={{ width: 132, flexShrink: 0 }} />
        <div style={{ flex: 1, position: 'relative', height: 18 }}>
          {ticks.map((t, i) => (
            <Text key={i} type="secondary" style={{ position: 'absolute', left: `${i * 25}%`, fontSize: 11, transform: i === 4 ? 'translateX(-100%)' : undefined }}>
              {dayjs(t).format('MM-DD HH:mm')}
            </Text>
          ))}
        </div>
        <div style={{ width: 96, flexShrink: 0 }} />
      </div>
      <div style={{ position: 'relative' }}>
        {showNow && (
          <div style={{ position: 'absolute', top: 0, bottom: 0, left: `calc(140px + (100% - 244px) * ${nowPct / 100})`, width: 1, background: '#ff4d4f', zIndex: 1, pointerEvents: 'none' }}>
            <span style={{ position: 'absolute', top: -2, left: 2, fontSize: 10, color: '#ff4d4f' }}>now</span>
          </div>
        )}
        {items.map((item, idx) => (
          <GanttRow
            key={item.id ?? idx}
            item={item}
            color={colors[item.sequence % colors.length]}
            rangeMin={range.min}
            rangeSpan={range.span}
            clampMinMs={planStartMs}
            adjustable={adjustable}
            onAdjust={onAdjust}
          />
        ))}
      </div>
      {adjustable && (
        <Text type="secondary" style={{ fontSize: 12, marginTop: 6, display: 'block' }}>
          拖拽条形调整开工时间（后续工单自动顺延）；拖拽中 ESC 取消；不可早于计划开始时间
        </Text>
      )}
    </div>
  );
};

const ApsPage: React.FC = () => {
  const [lines, setLines] = useState<ProductionLine[]>([]);
  const [linesError, setLinesError] = useState<string | null>(null);
  const [plans, setPlans] = useState<SchedulePlan[]>([]);
  const [plansLoading, setPlansLoading] = useState(false);
  const [selectedPlan, setSelectedPlan] = useState<SchedulePlan | null>(null);
  const [generating, setGenerating] = useState(false);
  const [releasingId, setReleasingId] = useState<string | null>(null);
  const [form] = Form.useForm();

  const loadLines = useCallback(async () => {
    setLinesError(null);
    try {
      const resp = await api('/api/mom/productionbase/productionline/all');
      setLines(resp?.data ?? []);
      if (!resp?.data?.length) setLinesError('暂无产线数据，请先在生产基础数据中维护产线');
    } catch (e) {
      setLinesError(String(e));
    }
  }, []);

  const loadPlans = useCallback(async () => {
    setPlansLoading(true);
    try {
      const resp = await api('/api/mom/aps/schedule/query?pageIndex=1&pageSize=50');
      setPlans(resp?.data ?? []);
    } catch (e) {
      message.error(String(e));
    } finally {
      setPlansLoading(false);
    }
  }, []);

  const loadPlanDetail = useCallback(async (id: string) => {
    try {
      const resp = await api(`/api/mom/aps/schedule/detail?id=${id}`);
      setSelectedPlan(resp?.data ?? null);
    } catch (e) {
      message.error(String(e));
    }
  }, []);

  useEffect(() => {
    loadLines();
    loadPlans();
  }, [loadLines, loadPlans]);

  const onGenerate = async (values: any) => {
    setGenerating(true);
    try {
      const resp = await api('/api/mom/aps/schedule/generate', {
        method: 'POST',
        body: JSON.stringify({
          productionLineID: values.productionLineID,
          startTime: values.startTime ? values.startTime.format(fmt) : '',
          dailyWorkMinutes: values.dailyWorkMinutes ?? 0,
          changeoverSeconds: values.changeoverSeconds ?? 0,
        }),
      });
      message.success(`排程计划已生成：${resp?.planNo ?? ''}`);
      await loadPlans();
      if (resp?.planID) await loadPlanDetail(resp.planID);
    } catch (e) {
      message.error(String(e), 5);
    } finally {
      setGenerating(false);
    }
  };

  const onRelease = async (planID: string) => {
    setReleasingId(planID);
    try {
      await api('/api/mom/aps/schedule/release', { method: 'PUT', body: JSON.stringify({ planID }) });
      message.success('排程已下发，工单预计开完工时间已回写');
      await loadPlans();
      if (selectedPlan?.id === planID) await loadPlanDetail(planID);
    } catch (e) {
      message.error(String(e));
    } finally {
      setReleasingId(null);
    }
  };

  const onAdjust = async (item: ScheduleItem, newStartMs: number) => {
    if (!selectedPlan) return;
    try {
      await api('/api/mom/aps/schedule/adjust', {
        method: 'PUT',
        body: JSON.stringify({
          planID: selectedPlan.id,
          itemID: item.id,
          newStartTime: dayjs(newStartMs).format(fmt),
          cascade: true,
        }),
      });
      message.success(`已调整 ${item.productOrderNo} 并顺延后续工单`);
      await loadPlanDetail(selectedPlan.id);
      await loadPlans();
    } catch (e) {
      message.error(String(e));
      await loadPlanDetail(selectedPlan.id);
    }
  };

  const columns = [
    { title: '批次号', dataIndex: 'planNo', width: 150 },
    { title: '产线', dataIndex: 'productionLineCode', width: 90 },
    { title: '工单数', dataIndex: 'orderCount', width: 80 },
    { title: '开始', width: 130, render: (_: any, r: SchedulePlan) => r.startTime ? dayjs(r.startTime).format('MM-DD HH:mm') : '-' },
    { title: '结束', width: 130, render: (_: any, r: SchedulePlan) => r.endTime ? dayjs(r.endTime).format('MM-DD HH:mm') : '-' },
    {
      title: '状态', dataIndex: 'currentState', width: 90,
      render: (v: string) => <Tag color={v === '已下发' ? 'green' : v === '已作废' ? 'default' : 'blue'}>{v}</Tag>,
    },
    {
      title: '操作', width: 170,
      render: (_: any, r: SchedulePlan) => (
        <Space>
          <Button size="small" onClick={() => loadPlanDetail(r.id)}>查看甘特</Button>
          {r.currentState === '已生成' && (
            <Button size="small" type="primary" loading={releasingId === r.id} onClick={() => onRelease(r.id)}>下发</Button>
          )}
        </Space>
      ),
    },
  ];

  return (
    <div style={{ padding: 16, background: '#f5f5f5', minHeight: '100%' }}>
      <Title level={4} style={{ marginTop: 0 }}>APS 排程</Title>

      <Card size="small" title="生成排程（优先级 → 交期，支持换型约束 / 物料齐套 / 按日窗口）">
        <Form form={form} layout="inline" onFinish={onGenerate} initialValues={{ dailyWorkMinutes: 0, changeoverSeconds: 0 }}>
          <Form.Item name="productionLineID" label="产线" rules={[{ required: true, message: '请选择产线' }]}>
            <Select
              style={{ width: 220 }} placeholder="选择产线" showSearch optionFilterProp="label"
              notFoundContent={linesError ? (
                <Space direction="vertical" size={4} style={{ padding: 8 }}>
                  <Text type="danger" style={{ fontSize: 12 }}>{linesError}</Text>
                  <Button size="small" icon={<ReloadOutlined />} onClick={loadLines}>重试</Button>
                </Space>
              ) : <Spin size="small" />}
              options={lines.map((l) => ({ value: l.id, label: `${l.code} ${l.description ?? ''}`.trim(), key: l.id }))}
            />
          </Form.Item>
          <Form.Item name="startTime" label="起始时间">
            <DatePicker showTime={{ format: 'HH:mm' }} format="YYYY-MM-DD HH:mm" style={{ width: 190 }} />
          </Form.Item>
          <Form.Item name="dailyWorkMinutes" label="每日工时(分,0=连续)" tooltip="按日窗口排程：跨天工时自动顺延到次日窗口">
            <InputNumber min={0} style={{ width: 155 }} />
          </Form.Item>
          <Form.Item name="changeoverSeconds" label="换型时间(秒)" tooltip="相邻工单产品型号切换时插入的换型间隔">
            <InputNumber min={0} style={{ width: 125 }} />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" loading={generating}>生成排程</Button>
          </Form.Item>
        </Form>
      </Card>

      <Card
        size="small" title="排程计划" style={{ marginTop: 12 }} styles={{ body: { padding: 0 } }}
        extra={<Button size="small" icon={<ReloadOutlined />} onClick={loadPlans} loading={plansLoading}>刷新</Button>}
      >
        <Table
          rowKey="id" size="small" columns={columns} dataSource={plans}
          loading={plansLoading} pagination={{ pageSize: 10, showSizeChanger: false }}
          locale={{ emptyText: '暂无排程计划，请在上方生成' }}
        />
      </Card>

      {selectedPlan && (
        <Card
          size="small" style={{ marginTop: 12 }}
          title={`甘特图 — ${selectedPlan.planNo}（${selectedPlan.productionLineCode || '-'}，${selectedPlan.orderCount} 个工单）`}
          extra={selectedPlan.currentState === '已生成'
            ? <Tag color="blue">可拖拽调整</Tag>
            : <Tag>{selectedPlan.currentState}（只读）</Tag>}
        >
          <Gantt
            items={selectedPlan.items ?? []}
            planStartMs={selectedPlan.startTime ? new Date(selectedPlan.startTime).getTime() : 0}
            adjustable={selectedPlan.currentState === '已生成'}
            onAdjust={onAdjust}
          />
        </Card>
      )}
    </div>
  );
};

export default ApsPage;
