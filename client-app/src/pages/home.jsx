import React from 'react';
import {
  Page,
  Navbar,
  NavTitle,
  NavTitleLarge,
  Link,
  Toolbar,
  Block,
} from 'framework7-react';
import Navigation from '../components/Navigation'
export default () => (
  <Page name="home">
    <Navigation></Navigation>
    {/* Top Navbar */}
    <Navbar large>
      <NavTitle>ch3ck1n</NavTitle>
      <NavTitleLarge>ch3ck1n</NavTitleLarge>
    </Navbar>
    {/* Toolbar */}
    <Toolbar bottom>
      <Link href="/Privacy/">Privacy</Link>
      <Link>Right Link</Link>
    </Toolbar>
    {/* Page content */}
    <Block strong>
      <p>Here is your blank Framework7 app. Let's see what we have here.</p>
    </Block>

  </Page>
);