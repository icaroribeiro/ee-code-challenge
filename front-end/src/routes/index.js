import React from 'react';
import { Switch } from 'react-router-dom';
import { ConnectedRouter } from 'connected-react-router';

import history from './history.js';
import DynamicLayoutRoute from '../components/DynamicLayoutRoute/index.js';
import Home from '../pages/Home/index.js';
import SubmitContainer from '../containers/SubmitContainer/index.js';
import Loading from '../components/Loading/index.js';
import NotFound from '../pages/NotFound/index.js';

const Routes = () => {
    return (
        <ConnectedRouter history={ history }>
            <Switch>           
                <DynamicLayoutRoute exact path="/" component={ Home } layout={ "NAV" } />
                <DynamicLayoutRoute path='/repositories' layout={ "NAV" } render={ (props) => (
                    <SubmitContainer {...props } placeholder={ Loading } />
                ) }/>
                <DynamicLayoutRoute path="*" component={ NotFound } />       
            </Switch>
        </ConnectedRouter>
    )
}

export default Routes;