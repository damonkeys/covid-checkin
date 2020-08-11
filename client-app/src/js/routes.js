
import HomePage from '../pages/home.jsx';
import Privacy from '../pages/Privacy';
import NoEmail from '../pages/NoEmail';
import NotFoundPage from '../pages/NotFoundPage';
import Activation from '../pages/Activation';

var routes = [
  {
    path: '/',
    component: HomePage,
  }, {
    path: '/privacy',
    component: Privacy,
  }, {
    path: '/NoEmail',
    component: NoEmail
  }, {
    path: '/activation/:activationToken',
    component: Activation
  }, {
    path: '(.*)',
    component: NotFoundPage
  },
];

export default routes;
