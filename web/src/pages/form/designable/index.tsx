import {
  Designer,
  DesignerToolsWidget,
  ViewToolsWidget,
  Workspace,
  OutlineTreeWidget,
  ResourceWidget,
  HistoryWidget,
  StudioPanel,
  CompositePanel,
  WorkspacePanel,
  ToolbarPanel,
  ViewportPanel,
  ViewPanel,
  SettingsPanel,
  ComponentTreeWidget
} from '@swiftease/designable-react'
import { SettingsForm, setNpmCDNRegistry } from '@swiftease/designable-react-settings-form'
import {
  createDesigner,
  Shortcut,
  KeyCode,
} from '@swiftease/designable-core'
import {
  ActionsWidget,
  SchemaEditorWidget,
  MarkupSchemaWidget,
  Inputs,
  AllDesignerComponents,
  Arrays, Displays, Layouts, CustomerComponents, saveSchema, funcs, commonComponents, chartComponents,AtaliValueInput,AtaliImageInput,bigScreenComponents
} from '@swiftease/atali-form'
import { BMapAddressPicker } from '@swiftease/atali-components'
import { PreviewWidget } from './PreviewWidget'
import { useParams } from '@umijs/max'
setNpmCDNRegistry('//unpkg.com')


const SaveShortCut = new Shortcut({
  codes: [
    [KeyCode.Meta, KeyCode.S],
    [KeyCode.Control, KeyCode.S],
  ],
  handler(ctx) {
    const list = window.location.href.split("/")
    saveSchema(ctx.engine, list.pop() || "", () => { })
  },
})

const engine = createDesigner({
  shortcuts: [
    SaveShortCut
  ],
  rootComponentName: 'Form',
})

export default (props: any) => {
  const params  = useParams();
  return (
    <Designer engine={engine}>
      <StudioPanel actions={<ActionsWidget formID={params?.formID} />}>
        <CompositePanel>
          <CompositePanel.Item title="panels.Component" icon="Component">
          <ResourceWidget title="sources.Common" sources={commonComponents} />
          <ResourceWidget title="sources.Charts" sources={chartComponents} />
            <ResourceWidget title="sources.Customer" sources={[...CustomerComponents, BMapAddressPicker]} />
            <ResourceWidget title="sources.Inputs" sources={Inputs} />
            <ResourceWidget title="sources.Layouts" sources={Layouts} />
            <ResourceWidget title="sources.Arrays" sources={Arrays} />
            <ResourceWidget title="sources.Displays" sources={Displays} />
            <ResourceWidget title="sources.BigScreens" sources={bigScreenComponents} />
          </CompositePanel.Item>
          <CompositePanel.Item title="panels.OutlinedTree" icon="Outline">
            <OutlineTreeWidget />
          </CompositePanel.Item>
          <CompositePanel.Item title="panels.History" icon="History">
            <HistoryWidget />
          </CompositePanel.Item>
        </CompositePanel>
        <Workspace id="form">
          <WorkspacePanel>
            <ToolbarPanel>
              <DesignerToolsWidget />
              <ViewToolsWidget
                use={['DESIGNABLE', 'JSONTREE', 'MARKUP', 'PREVIEW']}
              />
            </ToolbarPanel>
            <ViewportPanel>
              <ViewPanel type="DESIGNABLE">
                {() => (
                  <ComponentTreeWidget
                    components={{ ...AllDesignerComponents, BMapAddressPicker }}
                  />
                )}
              </ViewPanel>
              <ViewPanel type="JSONTREE" scrollable={false}>
                {(tree, onChange) => (
                  <SchemaEditorWidget tree={tree} onChange={onChange} />
                )}
              </ViewPanel>
              <ViewPanel type="MARKUP" scrollable={false}>
                {(tree) => <MarkupSchemaWidget tree={tree} />}
              </ViewPanel>
              <ViewPanel type="PREVIEW">
                {(tree) => <PreviewWidget tree={tree} />}
              </ViewPanel>
            </ViewportPanel>
          </WorkspacePanel>
        </Workspace>
        <SettingsPanel title="panels.PropertySettings">
          <SettingsForm scope={funcs} components={{AtaliValueInput,AtaliImageInput}} />
        </SettingsPanel>
      </StudioPanel>
    </Designer>
  )
}
