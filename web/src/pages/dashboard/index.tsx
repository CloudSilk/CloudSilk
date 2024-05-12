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
        <FormPreviewPage showPageContainer={false} title={''} showButton={false} createSchemaField={createSchemaField} formID={formID ?? "7877b188-2593-4c1c-bb1e-7ca7eb9dc0f5"}></FormPreviewPage>
    )
}

export default (props: any) => {
    return (
        <Suspense fallback={null}>
            <Dashboard {...props}></Dashboard>
        </Suspense>
    )
}