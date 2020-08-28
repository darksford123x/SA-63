import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'Record' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`Repair Slip ${profile.givenName || 'to Backstage'}`}
       subtitle="for information."
     ></Header>
     <Content>
       <ContentHeader title="add information">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             Add 
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
