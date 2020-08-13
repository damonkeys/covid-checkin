// @flow
import React, {Component} from 'react';
import {
    Button,
    Col,
    Link,
    Row
} from 'framework7-react';

type Props = {
    compact?: boolean,
    callbackURL?: string
}

type State = {
}

export default class Logins extends Component<Props, State> {
    $f7: any;

    render() {
        if (this.props.compact) {
            return(
                <Row>
                <Col width="10"></Col>
                <Col width="80" className="text-align-center">
                        <br />
                        {!this.$f7.device.ios && !this.$f7.device.macos ? (null) :
                            (
                                <Link color="white" style={{backgroundColor: '#000'}} className="socialmedia-button-compact" iconF7="logo_apple" href={'/auth/login?provider=apple&callbackUrl=' + (this.props.callbackURL || '/')} external ignoreCache />
                            )
                        }
                        <Link fill color="white" style={{backgroundColor: "#3b5998"}} className="socialmedia-button-compact" iconF7="logo_facebook" href={'/auth/login?provider=facebook&callbackUrl=' + (this.props.callbackURL || '/')} external ignoreCache />
                        <Link fill color="white" style={{backgroundColor: "#db4437"}} className="socialmedia-button-compact" iconF7="logo_googleplus" href={'/auth/login?provider=gplus&callbackUrl=' + (this.props.callbackURL || '/')} external ignoreCache />
                </Col>
                <Col width="10"></Col>
            </Row>
            );
        }

        return(
            <Row>
                <Col width="10"></Col>
                <Col width="80">
                    <br />
                    {!this.$f7.device.ios && !this.$f7.device.macos ? (null) :
                        (
                            <span><Button large fill color="black" iconF7="logo_apple" text="Apple" href={"/auth/login?provider=apple&callbackUrl=/"} external ignoreCache></Button><br /></span>
                        )
                    }
                    <Button large fill color="blue" style={{backgroundColor: "#3b5998"}} iconF7="logo_facebook" text="Facebook" href={"/auth/login?provider=facebook&callbackUrl=/"} external ignoreCache></Button><br />
                    <Button large fill color="pink" style={{backgroundColor: "#db4437"}} iconF7="logo_googleplus" text="Google" href={"/auth/login?provider=gplus&callbackUrl=/"} external ignoreCache></Button><br />
                </Col>
                <Col width="10"></Col>
            </Row>
        )
    }
}
