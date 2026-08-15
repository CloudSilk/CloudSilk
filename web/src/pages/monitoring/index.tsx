import React, { useCallback, useEffect, useRef, useState } from 'react';
import {
  Badge, Button, Card, Col, Row, Segmented, Space, Statistic, Table, Tag, Tooltip, Typography,
} from 'antd';
import { ReloadOutlined } from '@ant-design/icons';
import dayjs from 'dayjs';
import umiRequest from 'umi-request';
import { getToken } from '@swiftease/atali-pkg';

const { Title, Text } = Typography;

/** 后端统一响应码：Code_Success = 20000 */
const CODE_SUCCESS = 20000;

interface Overview {
  producingOrderCount: number;
  productionLineCount: number;
  productionStationCount: number;
  totalOrderQTY: number;
  totalFinishedQTY: number;
  totalStartedQTY: number;
  unhandledAlarmCount: number;
  todayAlarmCount: number;
  idleStationCount: number;
  producingStationCount: number;
  breakdownStationCount: number;
}

interface AlarmItem {
  id: string;
  alarmNo: string;
  alarmMessage: string;
  productionStationCode: string;
  currentState: string;
  createTime: string;
}

const authHeaders = (): Record<string, string> => ({
  'Content-Type': 'application/json',
  authorization: 'Bearer ' + getToken(),
});

const api = async (url: string): Promise<any> => {
  const resp = await umiRequest<Record<string, any>>(url, { method: 'GET', headers: authHeaders() });
  if (resp && (resp.code === CODE_SUCCESS || (resp.code === undefined && resp.data !== undefined))) {
    return resp;
  }
  throw new Error(resp?.message || `请求失败（code=${resp?.code}）`);
};

const REFRESH_OPTIONS = [
  { label: '5秒', value: 5 },
  { label: '30秒', value: 30 },
  { label: '关闭', value: 0 },
];

const StatCard: React.FC<{
  title: string;
  value: string | number;
  suffix?: string;
  tone?: 'default' | 'success' | 'warning' | 'danger';
  loading?: boolean;
}> = ({ title, value, suffix, tone, loading }) => {
  const color = tone === 'danger' ? '#cf1322' : tone === 'warning' ? '#d46b08'
    : tone === 'success' ? '#389e0d' : undefined;
  return (
    <Card size="small" loading={loading}>
      <Statistic
        title={<span style={{ fontSize: 12 }}>{title}</span>}
        value={value}
        suffix={suffix ? <span style={{ fontSize: 13 }}>{suffix}</span> : undefined}
        valueStyle={{ fontSize: 24, fontWeight: 600, color }}
      />
    </Card>
  );
};

const MonitoringPage: React.FC = () => {
  const [overview, setOverview] = useState<Overview | null>(null);
  const [alarms, setAlarms] = useState<AlarmItem[]>([]);
  const [alarmsLoading, setAlarmsLoading] = useState(true);
  const [firstLoading, setFirstLoading] = useState(true);
  const [lastUpdated, setLastUpdated] = useState<dayjs.Dayjs | null>(null);
  const [live, setLive] = useState<boolean | null>(null); // null=未知 true=正常 false=异常
  const [intervalSec, setIntervalSec] = useState<number>(5);
  const refreshLock = useRef(false);

  const fetchOverview = useCallback(async () => {
    const resp = await api('/api/mom/monitoring/overview');
    setOverview(resp.data ?? null);
  }, []);

  const fetchAlarms = useCallback(async () => {
    const resp = await api('/api/mom/monitoring/alarms?currentState=%E6%9C%AA%E5%A4%84%E7%90%86&limit=50');
    setAlarms(resp.data ?? []);
  }, []);

  const refresh = useCallback(async (withAlarms = true) => {
    if (refreshLock.current) return;
    refreshLock.current = true;
    try {
      await fetchOverview();
      if (withAlarms) await fetchAlarms();
      setLive(true);
      setLastUpdated(dayjs());
    } catch {
      setLive(false);
    } finally {
      refreshLock.current = false;
      setFirstLoading(false);
      setAlarmsLoading(false);
    }
  }, [fetchOverview, fetchAlarms]);

  useEffect(() => {
    refresh();
  }, [refresh]);

  useEffect(() => {
    if (!intervalSec) return;
    const t = setInterval(() => refresh(), intervalSec * 1000);
    return () => clearInterval(t);
  }, [intervalSec, refresh]);

  const o = overview;
  const finishRate = o && o.totalOrderQTY > 0 ? Math.round((o.totalFinishedQTY / o.totalOrderQTY) * 100) : null;

  const columns = [
    { title: '时间', dataIndex: 'createTime', width: 150, render: (v: string) => v ? dayjs(v).format('MM-DD HH:mm:ss') : '-' },
    { title: '工位', dataIndex: 'productionStationCode', width: 100 },
    { title: '报警编号', dataIndex: 'alarmNo', width: 140 },
    { title: '报警信息', dataIndex: 'alarmMessage', ellipsis: true },
    {
      title: '状态', dataIndex: 'currentState', width: 90,
      render: (v: string) => <Tag color={v === '未处理' ? 'red' : 'default'}>{v}</Tag>,
    },
  ];

  return (
    <div style={{ padding: 16, background: '#f5f5f5', minHeight: '100%' }}>
      <div style={{ marginBottom: 12, display: 'flex', justifyContent: 'space-between', alignItems: 'center', flexWrap: 'wrap', gap: 8 }}>
        <Title level={4} style={{ margin: 0 }}>厂级生产监控</Title>
        <Space size={12}>
          <Text type="secondary" style={{ fontSize: 12 }}>
            {lastUpdated ? `更新于 ${lastUpdated.format('HH:mm:ss')}` : '加载中…'}
          </Text>
          <Badge
            status={live === null ? 'default' : live ? 'processing' : 'error'}
            text={live === null ? '待首次加载' : live ? '数据正常' : '数据异常'}
          />
          <Segmented options={REFRESH_OPTIONS} value={intervalSec} onChange={(v) => setIntervalSec(v as number)} />
          <Button size="small" icon={<ReloadOutlined />} onClick={() => refresh()}>刷新</Button>
        </Space>
      </div>

      <Row gutter={[12, 12]}>
        <Col xs={12} sm={8} md={6} lg={3}>
          <StatCard title="生产中工单" value={o?.producingOrderCount ?? 0} loading={firstLoading} />
        </Col>
        <Col xs={12} sm={8} md={6} lg={3}>
          <StatCard title="产线 / 工站" value={`${o?.productionLineCount ?? 0} / ${o?.productionStationCount ?? 0}`} loading={firstLoading} />
        </Col>
        <Col xs={12} sm={8} md={6} lg={3}>
          <Tooltip title={o ? `完工 ${o.totalFinishedQTY} / 下单 ${o.totalOrderQTY}` : ''}>
            <div>
              <StatCard
                title="完工 / 下单"
                value={o ? `${o.totalFinishedQTY}/${o.totalOrderQTY}` : '-'}
                tone={finishRate !== null && finishRate >= 90 ? 'success' : 'default'}
                loading={firstLoading}
              />
              {finishRate !== null && (
                <Text type="secondary" style={{ fontSize: 11 }}>达成率 {finishRate}%</Text>
              )}
            </div>
          </Tooltip>
        </Col>
        <Col xs={12} sm={8} md={6} lg={3}>
          <StatCard title="未处理告警" value={o?.unhandledAlarmCount ?? 0} tone={(o?.unhandledAlarmCount ?? 0) > 0 ? 'danger' : 'success'} loading={firstLoading} />
        </Col>
        <Col xs={12} sm={8} md={6} lg={3}>
          <StatCard title="今日告警" value={o?.todayAlarmCount ?? 0} tone={(o?.todayAlarmCount ?? 0) > 0 ? 'warning' : 'default'} loading={firstLoading} />
        </Col>
        <Col xs={12} sm={8} md={6} lg={2}>
          <StatCard title="待机工位" value={o?.idleStationCount ?? 0} loading={firstLoading} />
        </Col>
        <Col xs={12} sm={8} md={6} lg={2}>
          <StatCard title="作业工位" value={o?.producingStationCount ?? 0} tone="success" loading={firstLoading} />
        </Col>
        <Col xs={12} sm={8} md={6} lg={2}>
          <StatCard title="故障工位" value={o?.breakdownStationCount ?? 0} tone={(o?.breakdownStationCount ?? 0) > 0 ? 'danger' : 'success'} loading={firstLoading} />
        </Col>
      </Row>

      <Card
        size="small" title="未处理告警（实时）" style={{ marginTop: 12 }} styles={{ body: { padding: 0 } }}
        extra={<Badge count={alarms.length} overflowCount={99} />}
      >
        <Table
          rowKey="id"
          size="small"
          columns={columns}
          dataSource={alarms}
          loading={alarmsLoading}
          pagination={false}
          scroll={{ y: 420 }}
          locale={{ emptyText: '暂无未处理告警' }}
        />
      </Card>
    </div>
  );
};

export default MonitoringPage;
