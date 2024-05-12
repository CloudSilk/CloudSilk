import { EditComponent } from '@swiftease/atali-curd'
import { createSchemaField } from '../../../form/field'

export default (props: any) => {
    return <EditComponent {...props} createSchemaField={createSchemaField}></EditComponent>
}