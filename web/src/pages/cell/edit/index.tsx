import { EditCompomemtProps, EditComponentState } from '@swiftease/atali-curd'
import { funcs,defaultCache, newService } from '@swiftease/atali-form'
import { ArrowLeftOutlined } from '@ant-design/icons'
import { PageContainer } from '@ant-design/pro-layout'
import { Button, Col, Row } from 'antd'
import { createSchemaField } from '../../form/field'
import { FormConsumer, FormProvider } from '@formily/react'
import { Graph } from '@antv/x6'
import { createForm } from '@formily/core'
import React from 'react'
import { history } from '@umijs/max';
import { CellCache, CollapseGroup } from '@swiftease/atali-graph'
import { Cell } from '@swiftease/atali-pkg'
import queryString from 'query-string';

interface EditCellComponentState extends EditComponentState {
    graph?: Graph
    cellCache?: CellCache
}

export class EditCellComponent extends React.Component<EditCompomemtProps, EditCellComponentState> {
    private container: HTMLDivElement | undefined

    constructor(props: any) {
        super(props)
        const locationState = history.location?.state
        if (locationState) {
            this.state = {
                ...locationState
            }
        }
    }
    refContainer = (container: HTMLDivElement) => {
        this.container = container
    }

    createGraph(){
        const graph = new Graph({
            container: this.container,
            grid: true,
            connecting: {
                allowBlank: true,
                highlight: true,
                snap: true,
                allowMulti: false,
                allowLoop: false,
                allowEdge: false,
                router: {
                    name: 'er',
                    args: {
                        offset: 'center',
                    },
                }
            }
        })
        graph.on('node:collapse', ({ node }: { node: CollapseGroup }) => {
            node.toggleCollapse()
            const collapsed = node.isCollapsed()
            const collapse = (parent: CollapseGroup) => {
              const cells = parent.getChildren()
              if (cells) {
                cells.forEach((cell) => {
                  if (collapsed) {
                    cell.hide()
                  } else {
                    cell.show()
                  }
          
                  if (cell instanceof CollapseGroup) {
                    if (!cell.isCollapsed()) {
                      collapse(cell)
                    }
                  }
                })
              }
            }
          
            collapse(node)
          })
          return graph
    }

    async componentDidMount() {
        const params = queryString.parse(history.location.search);
        const pageConfig = await defaultCache.getPageConfig(params['pageName']);
        if (pageConfig?.code !== 20000 || !pageConfig.data) return;
        let editFormSchema = {
            type: 'object',
            properties: {},
        };
        if (pageConfig.data.addFormID && pageConfig.data.addFormID != "") {
            const resp = await defaultCache.getFormConfig(pageConfig.data.addFormID)
            if (resp?.code == 20000) {
                const schema = JSON.parse(resp.data.schema);
                editFormSchema = schema.schema
            }
        }
        this.setState({
            id: params['id'],
            isAdd: params['isAdd'] == 'true',
            pageConfig: pageConfig.data,
            editFormSchema: editFormSchema,
            service: newService<any>(pageConfig.data.path != "" ? pageConfig.data.path : "curd/common/" + pageConfig.data.name),
        },async ()=>{
            const graph=this.createGraph()
            const cellCache= new CellCache()
            cellCache.init('','https://hhm.xw.life')
            if (this.state && !this.state.isAdd) {
                eval(this.state.pageConfig?.loadDetailBefore ?? '')
                const resp = await this.state.service.detail2({ pageName: this.state.pageConfig?.name, id: this.state.id })
                if (resp?.code === 20000) {
                    eval(this.state.pageConfig?.loadDetailAfter ?? '')
                    this.setState({
                        editForm: createForm({
                            initialValues: resp.data
                        }), graph: graph, cellCache: cellCache
                    })
                }
            } else {
                this.setState({ editForm: createForm(), graph: graph, cellCache: cellCache })
            }
        })
        
    }

    onSubmit(values: any) {
        const locationState = this.state
        eval(locationState.pageConfig?.submitBefore ?? '')
        let data: any = undefined
        if (locationState.pageConfig?.path != "") {
            data = values
        } else {
            data = { pageName: locationState.pageConfig?.name, data: values }
        }
        if (locationState && !locationState.isAdd) {
            locationState.service?.update(data).then((resp) => {
                eval(locationState.pageConfig?.submitAfter ?? '')
                this.onBack()
            })
        } else if (locationState && locationState.isAdd) {
            locationState.service?.add(data).then((resp) => {
                eval(locationState.pageConfig?.submitAfter ?? '')
                this.onBack()
            })
        }
    }

    updateCell(values: Cell) {
        try {
            this.state.cellCache?.registerCell(values, true)
            this.state.graph?.clearCells()
            const node = this.state.cellCache?.registerToGrapCell(values)
            if (node) {
                if (values.isEdge) {
                    let newEdge = this.state.cellCache?.newEdge(values.name, {
                        source: [170, 160],
                        target: [550, 160]
                    })
                    if (newEdge)
                        this.state.graph?.addEdge(newEdge)
                    newEdge = this.state.cellCache?.newEdge(values.name, {
                        source: [650, 60],
                        target: [650, 260]
                    })
                    if (newEdge)
                        this.state.graph?.addEdge(newEdge)
                } else {
                    const newNode = this.state.cellCache?.createNode(values.name, {
                        x: 100, y: 100
                    })
                    if (newNode){
                        this.state.cellCache?.setPorts(newNode)
                        this.state.graph?.addNode(newNode)
                    }
                        
                }

            }
        } catch (error) {
            console.log(error)
        }
    }

    onBack() {
        history.back()
    }

    render() {
        return <PageContainer className={'pageEdit'}
            header={{
                title: this.state?.pageConfig?.title + '-编辑',
                extra: [<Button type="text" onClick={this.onBack}><ArrowLeftOutlined /></Button>]
            }}
            footer={[
                <Button onClick={this.onBack}>取消</Button>,
                <Button onClick={
                    () => {
                        this.onSubmit(this.state.editForm?.values)
                    }
                } type="primary">
                    提交
                </Button>,
            ]}
        >
            <Row>
                <Col span={8}><div style={{
                    width: '100%',
                    height: '800px'
                }} ref={this.refContainer} id={"graph"} /></Col>
                <Col span={16}><div style={{ "backgroundColor": "white", "padding": "10px" }}>
                    {this.state?.editForm && <FormProvider labelCol={6} wrapperCol={12}
                        form={this.state.editForm}
                    >
                        {createSchemaField(this.state?.editFormSchema, funcs, false)}
                        <FormConsumer>
                            {(form) => {
                                this.updateCell(form.values)
                                return ""
                            }}
                        </FormConsumer>
                    </FormProvider>}</div></Col>
            </Row>
        </PageContainer>
    }
}
export default EditCellComponent