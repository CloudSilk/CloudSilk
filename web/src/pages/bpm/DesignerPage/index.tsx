import React from 'react'
import { Graph, Cell, Node } from '@antv/x6'
import '@antv/x6-react-shape'
import './index.less'
import { CellCache, CollapseGroup, Process, ProcessCell } from '@swiftease/atali-graph'
import { Menu, message } from 'antd'
import { Code, CommonService } from '@swiftease/atali-pkg'
import { createSchemaField } from '@/pages/form/field'
import { newService } from '@swiftease/atali-form'
import { DesignerPage } from '@swiftease/atali-graph'
import { useParams } from '@umijs/max'
import { useMatch } from '@umijs/max'
import { useRouteData } from '@umijs/max'
import { matchPath } from '@umijs/max'

interface BPMDesignerPageState {
    process?: Process
    processService: CommonService<Process>
    processID: string
}

// 画布操作菜单（撤销/重做/缩放/适应画布/清空）
const createCanvasMenu = (graph: Graph) => (
    <Menu onClick={(e) => {
        switch (e.key) {
            case 'undo':
                (graph as any).history?.undo()
                break
            case 'redo':
                (graph as any).history?.redo()
                break
            case 'zoomIn':
                graph.zoom(0.1)
                break
            case 'zoomOut':
                graph.zoom(-0.1)
                break
            case 'fit':
                graph.zoomToFit({ padding: 40 })
                break
            case 'clear':
                graph.clearCells()
                message.success('画布已清空')
                break
        }
    }} style={{ width: 200, height: '100%' }} mode="vertical">
        <Menu.Item key="undo">撤销</Menu.Item>
        <Menu.Item key="redo">重做</Menu.Item>
        <Menu.Divider />
        <Menu.Item key="zoomIn">放大</Menu.Item>
        <Menu.Item key="zoomOut">缩小</Menu.Item>
        <Menu.Item key="fit">适应画布</Menu.Item>
        <Menu.Divider />
        <Menu.Item key="clear" danger>清空画布</Menu.Item>
    </Menu>
)

export default class BPMDesignerPage extends React.Component<any, BPMDesignerPageState> {

    constructor(props: any) {
        super(props);
        const match = matchPath(
            { path: "/bpm/designer/:processID" },
            location.hash.replace("#",""),
          );
        const params  = match?.params;
        this.state = {
            processService: newService<Process>('bpm/process'),
            processID: params?.processID ?? ''
        }
    }

    cellToStep(self: BPMDesignerPage, cell: Cell.Properties, c: Cell): ProcessCell {
        let name = c.attr('text/text') as string
        if (c.isEdge()) {
            name = c.getLabelAt(0)?.attrs?.label?.text?.toString() ?? ''
        }

        const processCell = {
          name: name,
          id: cell.id,
          shape: cell.shape,
          zIndex: cell.zIndex,
          position: cell.position,
          x: cell.position?.x,
          y: cell.position?.y,
          data: c.getData(),
          sourceCell: cell.source?.cell,
          sourceSelector: cell.source?.selector,
          targetCell: cell.target?.cell,
          targetSelector: cell.target?.selector,
          width: cell.size?.width,
          height: cell.size?.height,
          children: cell.children?.join(","),
          size: undefined,
          source: undefined,
          target: undefined,
          originPositionX: cell.originPosition?.x,
          originPositionY: cell.originPosition?.y,
          originSizeWidth: cell.originSize?.width,
          originSizeHeight: cell.originSize?.height,
          visible: cell.visible ?? true,
          markup: JSON.stringify(cell.markup),
          attrs: JSON.stringify(cell.attrs)
        } as any as ProcessCell

        if (c instanceof CollapseGroup) {
            processCell.collapsed = c.isCollapsed()
            if (processCell.collapsed) {
                const size = c.getExpandSize()
                processCell.width = size.width
                processCell.height = size.height
            }
        }

        if (cell.parent) {
            processCell.parentID = cell.parent?.toString()
        }

        if (!processCell.data) {
            processCell.data = {
                id: processCell.id
            } as any
        } else if (processCell.data && processCell.data.id === "") {
            processCell.data.id = processCell.id
        }
        return processCell as any as ProcessCell
    }

    save(self: BPMDesignerPage, graph: Graph, cellCache: CellCache) {
        const process = self.state.process
        if (!process || !graph) return
        const list = graph.toJSON();
        const steps: ProcessCell[] = []
        list?.cells.forEach((cell) => {
            const c = graph?.getCellById(cell.id ?? '')
            if (cell.name === "CollapseGroup") {
            }
            steps.push(self.cellToStep(self, cell, c))
        })
        process.cells = steps

        self.state?.processService?.update(process).then(resp => {
            if (resp?.code != Code.Success) {
                message.error(resp?.message)
            } else {
                message.success("保存成功")
            }
        })
    }
    async getData(self: BPMDesignerPage, graph: Graph, cellCache: CellCache, collapse: (parent: CollapseGroup, hide: boolean) => void) {
        const resp = await self.state.processService.detail(self.state.processID)
        if (resp?.code == Code.Success) {
            self.setState({ process: resp.data })
            let cells: Cell.Properties[] = []
            resp.data.cells?.forEach(processCell => {
                const cell: Cell.Properties = {
                    id: processCell.id,
                    shape: processCell.shape,
                    zIndex: processCell.zIndex,
                    position: {
                        x: processCell.x,
                        y: processCell.y
                    },
                    source: {
                        cell: processCell.sourceCell,
                        selector: processCell.sourceSelector
                    },
                    originPosition: {
                        x: processCell.originPositionX,
                        y: processCell.originPositionY
                    },
                    originSize: {
                        width: processCell.originSizeWidth,
                        height: processCell.originSizeHeight
                    },
                    target: {
                        cell: processCell.targetCell,
                        selector: processCell.targetSelector
                    },
                    data: processCell.data,
                    size: {
                        width: processCell.width,
                        height: processCell.height
                    },
                    children: processCell.children?.split(","),
                    parent: processCell.parentID,

                }
                
                cellCache.setText(cell, processCell.name)
                cells.push(cell)
            })
            graph.fromJSON(cells)
            // 折叠节点
            resp.data.cells?.forEach(processCell => {
                let c = graph.getCellById(processCell.id ?? "")
                if (c instanceof CollapseGroup && processCell.collapsed) {
                    c.toggleCollapse(processCell.collapsed)
                    collapse(c, processCell.collapsed)
                }
            })
        }
    }
    addSucess(self: BPMDesignerPage, node: Node.Properties) {
        if (!node.data) node.data = {}
        node.data.processID = self.state.processID
    }
    render() {
        return <DesignerPage
            createMenu={(graph) => {
                return createCanvasMenu(graph)
            }}
            fileUrlPrefix={window.location.origin}
            system='BPM'
            data={this.state.process}
            graphFormID='3f2c8c66-c963-4c89-8152-be2de58655a2'
            isBPM={false}
            getData={async (grapth, cellCache, collapse) => {
                await this.getData(this, grapth, cellCache, collapse)
            }}
            save={(graph, cellCache) => {
                this.save(this, graph, cellCache)
            }}
            createSchemaField={createSchemaField}
            addSucess={(node) => {
                this.addSucess(this, node)
            }}
        ></DesignerPage>
    }
}