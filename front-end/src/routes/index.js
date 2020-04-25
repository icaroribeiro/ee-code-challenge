import React from 'react';
import { Switch } from 'react-router-dom';
import { ConnectedRouter } from 'connected-react-router';

import history from './history.js';
import DynamicLayoutRoute from '../components/DynamicLayoutRoute';
import Home from '../pages/Home';
import SubmitContainer from '../containers/SubmitContainer';
import Loading from '../components/Loading';
import NotFound from '../pages/NotFound';

const Routes = () => {
  return (
    <ConnectedRouter history={ history }>
      <Switch>
        <DynamicLayoutRoute exact path="/" component={ Home } layout={ "NAV" } />
        <DynamicLayoutRoute path='/repositories' layout={ "NAV" } render={ (props) => (
          <SubmitContainer { ...props } placeholder={ Loading } />
        ) }/>
        <DynamicLayoutRoute path="*" component={ NotFound } />
      </Switch>
    </ConnectedRouter>
  )
}

export default Routes;