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
                                const mockedData: string = `这是一段模拟的对话数据。本次会话传入了${messages.length}条消息`;
                                return new Response(mockedData);
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