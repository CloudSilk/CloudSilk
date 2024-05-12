import { createSchemaField } from '../field'
import { FormPreviewPage } from '@swiftease/atali-form';
import { useParams } from '@umijs/max';
import { useLocation } from '@umijs/max';
import queryString from 'query-string';


export default (props: any) => {
    const params  = useParams();
    const location = useLocation()
    const query = queryString.parse(location.search);
    const showButton = query?.showButton==='true';
    const showPageContainer = query?.showPageContainer==='true';
    const title = query?.pageTitle;
    return (
        <FormPreviewPage showPageContainer={showPageContainer} title={title??''} showButton={showButton} createSchemaField={createSchemaField} formID={params?.formID}></FormPreviewPage>
    )
}