import { newSchemaField } from '@swiftease/atali-form'
import {
    FormLayout
} from '@swiftease/formily-antd-v5'
import { BMapAddressPicker } from '@swiftease/atali-components'

export const SchemaField = newSchemaField({BMapAddressPicker})

export function createSchemaField(formSchema: any, funcs: any, horizontal: boolean) {
    if (horizontal) return <SchemaField schema={formSchema} scope={funcs} />
    return <FormLayout labelCol={6} wrapperCol={16}><SchemaField schema={formSchema} scope={funcs} /></FormLayout>
}
window['createSchemaField']=createSchemaField