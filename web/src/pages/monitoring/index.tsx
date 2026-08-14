import React, { useEffect, useRef, useState } from 'react';
import { Badge, Card, Col, Row, Spin, Table, Typography } from 'antd';
import umiRequest from 'umi-request';
import { getToken } from '@swiftease/atali-pkg';

const { Text, Title } = Typography;

interface Overview {
  producingOrder: number;
  productionLineCount: number;
  productionStationCount: number;
  totalOrderQTY: number;
  totalFinishedQTY: number;
  unhandledAlarms: number;
  todayAlarms: number;
  idleStations: number;
  producingStations: number;
  breakdownStations: number;
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

const OverviewCard: React.FC<{ title: string; value: React.ReactNode; suffix?: string }> = ({ title, value, suffix }) => (
  <Card size="small">
    <Text type="secondary" style={{ fontSize: 12 }}>{title}</Text>
    <Title level={3} style={{ margin: '4px 0 0' }}>{value}<span style={{ fontSize: 14, marginLeft: 4 }}>{suffix}</span></Title>
  </Card>
);

const MonitoringPage: React.FC = () => {
  const [overview, setOverview] = useState<Overview | null>(null);
  const [alarms, setAlarms] = useState<AlarmItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [live, setLive] = useState(false);
  const esRef = useRef<EventSource | null>(null);

  const fetchOverview = async () => {
    const resp = await umiRequest<Record<string, any>>('/api/mom/monitoring/overview', {
      method: 'GET',
      headers: authHeaders(),
    });
    if (resp?.code === 20000 || resp?.data) {
      setOverview(resp.data ?? resp);
    }
  };

  const fetchAlarms = async () => {
    const resp = await umiRequest<Record<string, any>>('/api/mom/monitoring/alarms?currentState=未处理&limit=50', {
      method: 'GET',
      headers: authHeaders(),
    });
    if (resp?.data) {
      setAlarms(resp.data as AlarmItem[]);
    }
  };

  useEffect(() => {
    (async () => {
      try {
        await Promise.all([fetchOverview(), fetchAlarms()]);
      } finally {
        setLoading(false);
      }
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // SSE 实时推送（token 无法附带在 EventSource 头里，退化为 5 秒轮询大屏接口）
  useEffect(() => {
    const timer = setInterval(async () => {
      try {
        await fetchOverview();
        setLive(true);
      } catch {
        setLive(false);
      }
    }, 5000);
    return () => clearInterval(timer);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const columns = [
    { title: '时间', dataIndex: 'createTime', width: 160 },
    { title: '工位', dataIndex: 'productionStationCode', width: 100 },
    { title: '报警编号', dataIndex: 'alarmNo', width: 140 },
    { title: '报警信息', dataIndex: 'alarmMessage' },
    {
      title: '状态', dataIndex: 'currentState', width: 90,
      render: (v: string) => <Badge status={v === '未处理' ? 'error' : 'default'} text={v} />,
    },
  ];

  if (loading) {
    return <div style={{ padding: 48, textAlign: 'center' }}><Spin size="large" /></div>;
  }

  const o = overview;
  return (
    <div style={{ padding: 16, background: '#f5f5f5', minHeight: '100%' }}>
      <div style={{ marginBottom: 12, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Title level={4} style={{ margin: 0 }}>厂级生产监控</Title>
        <Badge status={live ? 'processing' : 'warning'} text={live ? '实时刷新中（5s）' : '等待数据…'} />
      </div>

      <Row gutter={[12, 12]}>
        <Col xs={12} sm={8} md={6} lg={3}><OverviewCard title="生产中工单" value={(o as any)?.producingOrderCount ?? 0} /></Col>
        <Col xs={12} sm={8} md={6} lg={3}><OverviewCard title="产线 / 工站" value={`${o?.productionLineCount ?? 0}/${o?.productionStationCount ?? 0}`} /></Col>
        <Col xs={12} sm={8} md={6} lg={3}><OverviewCard title="完工 / 下单" value={`${o?.totalFinishedQTY ?? 0}/${o?.totalOrderQTY ?? 0}`} /></Col>
        <Col xs={12} sm={8} md={6} lg={3}><OverviewCard title="未处理告警" value={(o as any)?.unhandledAlarmCount ?? 0} /></Col>
        <Col xs={12} sm={8} md={6} lg={3}><OverviewCard title="今日告警" value={(o as any)?.todayAlarmCount ?? 0} /></Col>
        <Col xs={12} sm={8} md={6} lg={2}><OverviewCard title="待机" value={(o as any)?.idleStationCount ?? 0} /></Col>
        <Col xs={12} sm={8} md={6} lg={2}><OverviewCard title="作业" value={(o as any)?.producingStationCount ?? 0} /></Col>
        <Col xs={12} sm={8} md={6} lg={2}><OverviewCard title="故障" value={(o as any)?.breakdownStationCount ?? 0} /></Col>
      </Row>

      <Card size="small" title="未处理告警（实时）" style={{ marginTop: 12 }} styles={{ body: { padding: 0 } }}>
        <Table
          rowKey="id"
          size="small"
          columns={columns}
          dataSource={alarms}
          pagination={false}
          scroll={{ y: 420 }}
          locale={{ emptyText: '暂无未处理告警' }}
        />
      </Card>
    </div>
  );
};

export default MonitoringPage;
