import React from "react";
import { Route } from 'react-router-dom';

import LayoutNav from "../LayoutNav";

const DynamicLayoutRoute = (props) => {
  const actualRouteComponent = (
    <Route
      { ...props }
    />
  );

  switch (props.layout) {
    case 'NAV': {
      return (
        <LayoutNav>
          { actualRouteComponent }
        </LayoutNav>
      )
    }
    default: {
      return (
          actualRouteComponent
      )
    }
  }
};

export default DynamicLayoutRoute;