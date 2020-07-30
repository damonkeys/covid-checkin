
import HomePage from '../pages/home.jsx';
import Privacy from '../pages/Privacy';
import NotFoundPage from '../pages/NotFoundPage';

var routes = [
  {
    path: '/',
    component: HomePage,
  }, {
    path: '/privacy',
    component: Privacy,
  }, {
    path: '(.*)',
    component: NotFoundPage
  },
];

export default routes;
