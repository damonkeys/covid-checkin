// @flow
import React, { Component } from 'react';
import { Navbar, NavLeft, NavTitle, Searchbar } from 'framework7-react';
import i18n from '../../components/i18n.js';

type Props = {
    hideBacklink?: boolean,
    searchContainer?: string,
    searchIn?: string,
    hideMenu?: boolean,
    className?: string,
}

type State = {
    hideBacklink: boolean,
    useronline: boolean,
    username: string
}

/**
 * Navigation component shows the navigation at the top of a site. It will be visible if a user is logged in only.//#endregion
 * The component connects the backend to give the userdata of a online user. After connecting the backend, the callbackSession prop
 * is called to use usersession-data in other components.
 * 
 * There are some different props:
 * 
 * - callbackSession: it is a callback-function to return the read usersession-data.
 * - hideBacklink: boolean true/false to show or hide the back-link in the navigation bar
 */
export default class Navigation extends Component<Props, State> {
    $f7: any

    constructor(props: Props) {
        super(props);
        this.state = {
            hideBacklink: true,
            useronline: false,
            username: ''
        };
    }

    render() {
        return (
            <Navbar color="pink" className={this.props.className || 'navbar-main'}>

                <NavLeft>
                    {(this.props.hideBacklink || this.$f7.views.main.router.history.length <= 1) ? (null) :
                        (
                            <NavLeft backLink={i18n.t('basic.back')}></NavLeft>
                        )
                    }
                </NavLeft>

                {this.props.searchContainer ? (
                    <Searchbar searchContainer={this.props.searchContainer} searchIn={this.props.searchIn} disableButtonText={""}></Searchbar>
                ) : (
                        <NavTitle>{i18n.t('basic.appname')}</NavTitle>
                    )}
            </Navbar>
        );
    }
}
