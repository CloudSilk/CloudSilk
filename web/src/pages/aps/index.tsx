import React, { useEffect, useMemo, useRef, useState } from 'react';
import { Alert, Button, Card, DatePicker, Form, InputNumber, message, Select, Space, Table, Tooltip, Typography } from 'antd';
import dayjs, { Dayjs } from 'dayjs';
import umiRequest from 'umi-request';
import { getToken } from '@swiftease/atali-pkg';

const { Title, Text } = Typography;

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

const api = async <T,>(url: string, options: Record<string, any> = {}): Promise<T | undefined> => {
  const resp = await umiRequest<Record<string, any>>(url, { headers: authHeaders(), ...options });
  if (resp?.code !== undefined && resp.code !== 20000 && resp.code !== 1 && !resp.data && resp.code !== 0) {
    throw new Error(resp.message || `请求失败（code=${resp.code}）`);
  }
  return resp as T;
};

/**
 * 甘特条：按整体时间轴比例渲染；条可水平拖拽，松手调用 onAdjust 落库
 * （像素位移按容器宽度换算为时间偏移，拖拽中乐观平移视觉反馈）
 */
const GanttBar: React.FC<{
  items: ScheduleItem[];
  onAdjust?: (itemID: string, newStartMs: number) => void;
  adjustable?: boolean;
}> = ({ items, onAdjust, adjustable }) => {
  const range = useMemo(() => {
    if (!items?.length) return null;
    const starts = items.map((i) => new Date(i.plannedStartTime).getTime()).filter((t) => !isNaN(t));
    const ends = items.map((i) => new Date(i.plannedEndTime).getTime()).filter((t) => !isNaN(t));
    if (!starts.length || !ends.length) return null;
    const min = Math.min(...starts);
    const max = Math.max(...ends);
    return { min, span: Math.max(max - min, 1) };
  }, [items]);

  const trackRefs = useRef<Record<string, HTMLDivElement | null>>({});
  const [dragging, setDragging] = useState<{ id: string; x0: number; dx: number; start: number } | null>(null);

  // 拖拽中：全局监听移动/抬起
  useEffect(() => {
    if (!dragging) return;
    const onMove = (e: MouseEvent) => {
      setDragging((d) => (d ? { ...d, dx: e.clientX - d.x0 } : d));
    };
    const onUp = (e: MouseEvent) => {
      const d = dragging;
      setDragging(null);
      if (d && onAdjust && range) {
        const track = trackRefs.current[d.id];
        if (track && track.clientWidth > 0) {
          const px = e.clientX - d.x0;
          const ms = (px / track.clientWidth) * range.span;
          onAdjust(d.id, d.start + ms);
        }
      }
    };
    window.addEventListener('mousemove', onMove);
    window.addEventListener('mouseup', onUp);
    return () => {
      window.removeEventListener('mousemove', onMove);
      window.removeEventListener('mouseup', onUp);
    };
  }, [dragging, onAdjust, range]);

  if (!range) return <Text type="secondary">暂无排程明细</Text>;

  const colors = ['#1677ff', '#52c41a', '#fa8c16', '#722ed1', '#13c2c2', '#eb2f96'];
  return (
    <div>
      <div style={dragging ? { userSelect: 'none' } : undefined}>
      {items.map((item, idx) => {
        const start = new Date(item.plannedStartTime).getTime();
        const end = new Date(item.plannedEndTime).getTime();
        if (isNaN(start) || isNaN(end)) return null;
        const itemKey = item.id ?? String(idx);
        const isDragging = dragging?.id === itemKey;
        const offsetPx = isDragging ? dragging!.dx : 0;
        const left = ((start - range.min) / range.span) * 100;
        const width = Math.max(((end - start) / range.span) * 100, 0.5);
        const color = colors[item.sequence % colors.length];
        return (
          <div key={itemKey} style={{ display: 'flex', alignItems: 'center', marginBottom: 4, gap: 8 }}>
            <Text style={{ width: 130, fontSize: 12 }} ellipsis>{`${item.sequence}. ${item.productOrderNo}`}</Text>
            <div ref={(el) => { trackRefs.current[itemKey] = el; }}
                 style={{ flex: 1, position: 'relative', height: 22, background: '#f0f0f0', borderRadius: 4 }}>
              <Tooltip title={`${item.plannedStartTime} → ${item.plannedEndTime}（${Math.round(item.durationSeconds / 60)} 分钟）${adjustable ? '，可拖拽调整' : ''}`}>
                <div
                  onMouseDown={(e) => {
                    if (!adjustable || !onAdjust) return;
                    setDragging({ id: itemKey, x0: e.clientX, dx: 0, start });
                  }}
                  style={{
                    position: 'absolute',
                    left: `calc(${left}% + ${offsetPx}px)`,
                    width: `${width}%`, top: 3, bottom: 3,
                    background: color, borderRadius: 3, minWidth: 6,
                    cursor: adjustable ? 'ew-resize' : 'pointer',
                    opacity: isDragging ? 0.75 : 1,
                    boxShadow: isDragging ? '0 0 0 2px rgba(22,119,255,.35)' : undefined,
                  }} />
              </Tooltip>
            </div>
          </div>
        );
      })}
      </div>
      <div style={{ display: 'flex', justifyContent: 'space-between', marginTop: 4 }}>
        <Text type="secondary" style={{ fontSize: 12 }}>{dayjs(range.min).format('MM-DD HH:mm')}</Text>
        <Text type="secondary" style={{ fontSize: 12 }}>{dayjs(range.min + range.span).format('MM-DD HH:mm')}</Text>
      </div>
    </div>
  );
};

const ApsPage: React.FC = () => {
  const [lines, setLines] = useState<ProductionLine[]>([]);
  const [plans, setPlans] = useState<SchedulePlan[]>([]);
  const [selectedPlan, setSelectedPlan] = useState<SchedulePlan | null>(null);
  const [generating, setGenerating] = useState(false);
  const [form] = Form.useForm();

  const loadLines = async () => {
    try {
      const resp = await api<{ data: ProductionLine[] }>('/api/mom/production/productionline/all');
      setLines((resp?.data as any) ?? []);
    } catch (e) {
      // 产线接口路径可能因部署差异不可用，允许手工输入场景由低代码页面承担
    }
  };

  const loadPlans = async () => {
    try {
      const resp = await api<{ data: SchedulePlan[] }>('/api/mom/aps/schedule/query?pageIndex=1&pageSize=50');
      setPlans((resp?.data as any) ?? []);
    } catch (e) {
      message.error(String(e));
    }
  };

  const loadPlanDetail = async (id: string) => {
    const resp = await api<{ data: SchedulePlan }>(`/api/mom/aps/schedule/detail?id=${id}`);
    setSelectedPlan((resp?.data as any) ?? null);
  };

  useEffect(() => {
    loadLines();
    loadPlans();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const onGenerate = async (values: any) => {
    if (!values.productionLineID) {
      message.warning('请选择产线');
      return;
    }
    setGenerating(true);
    try {
      const body = {
        productionLineID: values.productionLineID,
        startTime: values.startTime ? values.startTime.format('YYYY-MM-DD HH:mm:ss') : '',
        dailyWorkMinutes: values.dailyWorkMinutes ?? 0,
        changeoverSeconds: values.changeoverSeconds ?? 0,
      };
      const resp = await api<{ data: { planNo: string } }>('/api/mom/aps/schedule/generate', { method: 'POST', body: JSON.stringify(body) });
      message.success(`排程计划已生成：${(resp as any)?.planNo ?? (resp as any)?.data?.planNo ?? ''}`);
      await loadPlans();
    } catch (e) {
      message.error(String(e));
    } finally {
      setGenerating(false);
    }
  };

  const onAdjustItem = async (itemID: string, newStartMs: number) => {
    if (!selectedPlan) return;
    try {
      await api('/api/mom/aps/schedule/adjust', {
        method: 'PUT',
        body: JSON.stringify({
          planID: selectedPlan.id,
          itemID,
          newStartTime: dayjs(newStartMs).format('YYYY-MM-DD HH:mm:ss'),
          cascade: true,
        }),
      });
      message.success('已调整并顺延后续工单');
      await loadPlanDetail(selectedPlan.id);
      await loadPlans();
    } catch (e) {
      message.error(String(e));
      await loadPlanDetail(selectedPlan.id);
    }
  };

  const onRelease = async (planID: string) => {
    try {
      await api('/api/mom/aps/schedule/release', { method: 'PUT', body: JSON.stringify({ planID }) });
      message.success('排程已下发，工单预计开完工时间已回写');
      await loadPlans();
      if (selectedPlan?.id === planID) await loadPlanDetail(planID);
    } catch (e) {
      message.error(String(e));
    }
  };

  const columns = [
    { title: '批次号', dataIndex: 'planNo', width: 140 },
    { title: '产线', dataIndex: 'productionLineCode', width: 90 },
    { title: '工单数', dataIndex: 'orderCount', width: 80 },
    { title: '开始', dataIndex: 'startTime', width: 150, render: (v: string) => v ? dayjs(v).format('MM-DD HH:mm') : '-' },
    { title: '结束', dataIndex: 'endTime', width: 150, render: (v: string) => v ? dayjs(v).format('MM-DD HH:mm') : '-' },
    { title: '状态', dataIndex: 'currentState', width: 90 },
    {
      title: '操作', width: 170,
      render: (_: any, record: SchedulePlan) => (
        <Space>
          <Button size="small" onClick={() => loadPlanDetail(record.id)}>查看甘特</Button>
          {record.currentState === '已生成' && (
            <Button size="small" type="primary" onClick={() => onRelease(record.id)}>下发</Button>
          )}
        </Space>
      ),
    },
  ];

  return (
    <div style={{ padding: 16, background: '#f5f5f5', minHeight: '100%' }}>
      <Title level={4} style={{ marginTop: 0 }}>APS 排程</Title>

      <Card size="small" title="生成排程（前向贪心：优先级 → 交期，支持换型约束/物料齐套/按日窗口）">
        <Form form={form} layout="inline" onFinish={onGenerate} initialValues={{ dailyWorkMinutes: 0, changeoverSeconds: 0 }}>
          <Form.Item name="productionLineID" label="产线" rules={[{ required: true }]}>
            <Select style={{ width: 200 }} placeholder="选择产线" showSearch optionFilterProp="label"
              options={lines.map((l) => ({ value: l.id, label: `${l.code} ${l.description ?? ''}`, key: l.id }))} />
          </Form.Item>
          <Form.Item name="startTime" label="起始时间">
            <DatePicker showTime={{ format: 'HH:mm' }} format="YYYY-MM-DD HH:mm" style={{ width: 190 }} />
          </Form.Item>
          <Form.Item name="dailyWorkMinutes" label="每日工时(分,0=连续)">
            <InputNumber min={0} style={{ width: 150 }} />
          </Form.Item>
          <Form.Item name="changeoverSeconds" label="换型时间(秒)">
            <InputNumber min={0} style={{ width: 120 }} />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" loading={generating}>生成排程</Button>
          </Form.Item>
        </Form>
      </Card>

      <Card size="small" title="排程计划" style={{ marginTop: 12 }} styles={{ body: { padding: 0 } }}>
        <Table rowKey="id" size="small" columns={columns} dataSource={plans} pagination={{ pageSize: 10 }} />
      </Card>

      {selectedPlan && (
        <Card size="small" title={`甘特图 — ${selectedPlan.planNo}（${selectedPlan.productionLineCode}，${selectedPlan.orderCount} 个工单）`} style={{ marginTop: 12 }}
          extra={selectedPlan.currentState === '已生成' ? <Text type="secondary" style={{ fontSize: 12 }}>拖拽条形调整开工时间，后续工单自动顺延</Text> : undefined}>
          {selectedPlan.items?.length
            ? <GanttBar items={selectedPlan.items} onAdjust={onAdjustItem} adjustable={selectedPlan.currentState === '已生成'} />
            : <Alert type="info" message="该计划无明细" showIcon />}
        </Card>
      )}
    </div>
  );
};

export default ApsPage;
