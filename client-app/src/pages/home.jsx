import React, { Component } from 'react';
import {
  Page,
  Navbar,
  NavTitle,
  NavTitleLarge,
  Link,
  Toolbar,
  Block,
  Row,
  Col,
  Button,
} from 'framework7-react';

export default class Home extends Component<Props, State> {
  $f7router: any
  $f7route: any

  render() {
    return(
        <Page name="home">
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
          <p>Here is your blank Framework7 app. Let's see what we have here. Let's try some Logins!</p>
        </Block>
    
        <Row>
          <Col width="5">
          </Col>
          <Col>
              <Button large fill iconF7="logo_facebook" text="Facebook" href={"/auth/login?provider=facebook&callbackUrl=" + (this.$f7route.query.callbackUrl || '/')} external></Button><br />
              <Button large fill iconF7="logo_googleplus" text="Google" href={"/auth/login?provider=gplus&callbackUrl=" + (this.$f7route.query.callbackUrl || '/')} external></Button><br />
              <Button large fill iconF7="logo_apple" text="Apple" href={"/auth/login?provider=apple&callbackUrl=" + (this.$f7route.query.callbackUrl || '/')} external></Button>
          </Col>
          <Col width="5">
          </Col>
        </Row>
      </Page>
    );
  }    
}