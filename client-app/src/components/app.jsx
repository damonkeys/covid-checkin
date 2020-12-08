import React, { Suspense } from 'react';

import {
  App,
  View,
} from 'framework7-react';
import routes from '../js/routes';
import { I18nextProvider, Translation } from 'react-i18next';
import i18n from './i18n'; // the initialized i18next instance
import {
  Preloader, Block
} from 'framework7-react';

export default class extends React.Component {
  constructor() {
    super();

    this.state = {
      // Framework7 Parameters
      f7params: {
        name: 'chckr', // App name
        theme: 'auto', // Automatic theme detection



        // App routes
        routes: routes,
        // Register service worker
        //
        // IMPORTANT COMMENT:
        // service-worker.js uses precacheAndRoute. We have to configure it with URL outside of the app.
        // With service-worker.js included we now have the problem to a not working app. External chckr-calls
        // like connecting the auth-server doesn't work.
        // For more informations: https://developers.google.com/web/tools/workbox/modules/workbox-precaching#clean_urls 
        //
        // serviceWorker: {
        //   path: '/service-worker.js',
        // },
      },

    }
  }
  render() {
    return (
      <App params={ this.state.f7params } >
        {/* Your main view, should have "view-main" class */}
        <Suspense fallback={<Block className="text-align-center"><Preloader color="pink"></Preloader></Block>}>
          <I18nextProvider i18n={i18n}>
            <Translation>
              {
                t =>
                  <View main className="safe-areas"/>
              }
            </Translation>
          </I18nextProvider>
        </Suspense>
      </App>
    );
  }

  componentDidMount() {
    this.$f7ready((f7) => {

      // Call F7 APIs here
    });
  }
}
