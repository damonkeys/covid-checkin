
import HomePage from '../pages/home.jsx';
import Privacy from '../pages/Privacy';
import NoEmail from '../pages/NoEmail';
import ActivationMissing from '../pages/Activation/missing';
import ActivationStart from '../pages/Activation/start';
import NotFoundPage from '../pages/NotFoundPage';

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
    path: '/activation/missing',
    component: ActivationMissing
  }, {
    path: '/activation/start',
    component: ActivationStart
  }, {
    path: '(.*)',
    component: NotFoundPage
  },
];

export default routes;
