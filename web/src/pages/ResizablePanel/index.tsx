import { Suspense } from 'react'
import { ProChat } from '@ant-design/pro-chat';
import {
    getPanelElement,
    getPanelGroupElement,
    getResizeHandleElement,
    Panel,
    PanelGroup,
    PanelResizeHandle,
} from "react-resizable-panels";
import "./index.less"
import ResizeHandle from './ResizeHandle'

export default (props: any) => {
    return (
        <Suspense fallback={null}>
            <PanelGroup direction="horizontal" style={{margin: 5}}>
                <Panel className='Panel' defaultSize={8} maxSize={8} minSize={8}>
                    <div className='PanelContent'>left</div>
                </Panel>
                <Panel minSize={50}>
                    <div style={{ height: document.body.offsetHeight - 70,  backgroundColor: "white" }}>
                        <ProChat
                            helloMessage={
                                '欢迎使用 ProChat ，我是你的专属机器人，这是我们的 Github：[ProChat](https://github.com/ant-design/pro-chat)'
                            }
                            request={async (messages) => {
                                // 对接 OpenAI 兼容接口（usercenter AI 网关 /v1/chat/completions）
                                const resp = await fetch((process.env.WEB_BASE ?? '') + '/v1/chat/completions', {
                                    method: 'POST',
                                    headers: {
                                        'Content-Type': 'application/json',
                                        'Authorization': `Bearer ${localStorage.getItem('token') ?? ''}`,
                                    },
                                    body: JSON.stringify({
                                        model: 'default',
                                        messages: messages.map((m: any) => ({ role: m.role, content: m.content })),
                                        stream: false,
                                    }),
                                })
                                if (!resp.ok) {
                                    return new Response(`AI 服务请求失败（${resp.status}），请确认已部署 usercenter AI 网关。`)
                                }
                                const data = await resp.json()
                                return new Response(data?.choices?.[0]?.message?.content ?? '（空响应）')
                            }}
                        /></div>
                </Panel>
                <Panel defaultSize={15} maxSize={15} minSize={15}>
                    <PanelGroup direction="vertical">
                        <Panel defaultSize={30} minSize={20}>
                            <div className='PanelContent'>top</div>
                        </Panel>
                        <ResizeHandle />
                        <Panel minSize={30}>
                            <div className='PanelContent'>middle</div>
                        </Panel>
                        <ResizeHandle />
                        <Panel defaultSize={30} minSize={20}>
                            <div className='PanelContent'>bottom</div>
                        </Panel>
                    </PanelGroup>
                </Panel>
            </PanelGroup>

        </Suspense>
    )
}