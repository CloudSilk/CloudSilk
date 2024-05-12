import { useRef } from 'react';
import { ActionType } from '@ant-design/pro-table';
import { useLocation } from '@umijs/max';

import { TableCurdPage, TreeCurdPage, ListCurdPage, CardCurdPage, ProListCurdPage } from '@swiftease/atali-curd';
import { createSchemaField } from '../../form/field'
import { SplitCurdPage } from '@swiftease/atali-curd';
import queryString from 'query-string';
import { TabManager, TabContext } from '../../../components/TabManager';
import { uuid } from 'uuidv4'
export default () => {
    const location = useLocation()
    const actionRef = useRef<ActionType>()
    const query = queryString.parse(location.search);
    const pageName = query?.pageName ?? '';
    const pageType = Number(query?.type ?? 1);
    window["AtaliNoCache"]=true;
    // return (
    //     <>
    //         {pageType === 1 && <TableCurdPage createSchemaField={createSchemaField} pageName={pageName} actionRef={actionRef}></TableCurdPage>}
    //         {pageType === 2 && <TreeCurdPage createSchemaField={createSchemaField} pageName={pageName}></TreeCurdPage>}
    //         {pageType === 3 && <ListCurdPage createSchemaField={createSchemaField} pageName={pageName}></ListCurdPage>}
    //         {pageType === 4 && <CardCurdPage createSchemaField={createSchemaField} pageName={pageName}></CardCurdPage>}
    //         {pageType === 5 && <ProListCurdPage createSchemaField={createSchemaField} pageName={pageName} actionRef={actionRef}></ProListCurdPage>}
    //         {pageType === 6 && <SplitCurdPage createSchemaField={createSchemaField} pageName={pageName} actionRef={actionRef}></SplitCurdPage>}</>
    // );

    return <TabManager key={uuid()}>
            <>
                {pageType === 1 && <TableCurdPage  key={uuid()} createSchemaField={createSchemaField} pageName={pageName} actionRef={actionRef}></TableCurdPage>}
                {pageType === 2 && <TreeCurdPage  key={uuid()} createSchemaField={createSchemaField} pageName={pageName}></TreeCurdPage>}
                {pageType === 3 && <ListCurdPage  key={uuid()} createSchemaField={createSchemaField} pageName={pageName}></ListCurdPage>}
                {pageType === 4 && <CardCurdPage  key={uuid()} createSchemaField={createSchemaField} pageName={pageName}></CardCurdPage>}
                {pageType === 5 && <ProListCurdPage  key={uuid()} createSchemaField={createSchemaField} pageName={pageName} actionRef={actionRef}></ProListCurdPage>}
                {pageType === 6 && <SplitCurdPage  key={uuid()} createSchemaField={createSchemaField} pageName={pageName} actionRef={actionRef}></SplitCurdPage>}
            </>
        </TabManager>

};

