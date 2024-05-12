// 全局共享数据示例
import { DEFAULT_NAME } from '@/constants';
import { useState } from 'react';

const global = () => {
  const [name, setName] = useState<string>(DEFAULT_NAME);
  const [tabs, setTabs] = useState<any[]>([]);
  return {
    name,
    setName,
    tabs,
    setTabs
  };
};

export default global;
