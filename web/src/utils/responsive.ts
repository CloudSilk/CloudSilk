import { UAParser } from 'ua-parser-js';

/**
 * check mobile device in server
 */
export const isMobileDevice = () => {
  if (typeof process === 'undefined') {
    throw new Error('[Server method] you are importing a server-only module outside of server');
  }
  const ua = navigator.userAgent;
  
  console.log(ua);
  const device = new UAParser(ua || '').getDevice();

  return device.type === 'mobile';
};
