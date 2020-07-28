import React from 'react';

import {
  App,
  View,
} from 'framework7-react';

import routes from '../js/routes';

export default class extends React.Component {
  constructor() {
    super();

    this.state = {
      // Framework7 Parameters
      f7params: {
        name: 'ch3ck1n', // App name
        theme: 'auto', // Automatic theme detection



        // App routes
        routes: routes,
        // Register service worker
        serviceWorker: {
          path: '/service-worker.js',
        },
      },

    }
  }
  render() {
    return (
      <App params={ this.state.f7params } >

        {/* Your main view, should have "view-main" class */}
        <View main className="safe-areas" url="/" />

      </App>
    );
  }

  componentDidMount() {
    this.$f7ready((f7) => {

      // Call F7 APIs here
    });
  }
}