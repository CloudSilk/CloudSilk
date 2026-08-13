import { FormPreviewPage } from '@swiftease/atali-form'
import { createSchemaField } from '../form/field'
import { Suspense } from 'react'
import { useLocation } from '@umijs/max'
import queryString from 'query-string';


export function Dashboard(props: any) {
    const location = useLocation()
    const query = queryString.parse(location.search);
    const formID = query?.formID;
    return (
        <FormPreviewPage showPageContainer={false} title={''} showButton={false} createSchemaField={createSchemaField} formID={String(formID ?? process.env.DEFAULT_FORM_ID ?? '')}></FormPreviewPage>
    )
}

export default (props: any) => {
    return (
        <Suspense fallback={null}>
            <Dashboard {...props}></Dashboard>
        </Suspense>
    )
}