// @flow
import React, { Component } from 'react';
import {
    Block,
    Button,
    Input
} from 'framework7-react';
import i18n from '../../components/i18n.js';

type Props = {
}

type State = {
    locationCode: string
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
export default class CheckinInput extends Component<Props, State> {
    constructor(props: Props) {
        super(props);
        this.state = {
            locationCode: ''
        };
    }

    render() {
        return (
            <Block strong>
                {i18n.t('dashboard.explanation')}
                <h3 className="padding-large">{i18n.t('dashboard.location-code')}</h3>
                <Input id="location-code" label={i18n.t('dashboard.location-code')} type="text" placeholder={i18n.t('dashboard.location-code-placeholder')} onInput={(value) => {this.setState({locationCode: value.currentTarget.value});}}></Input>
                <Button rasied fill iconF7="checkmark" className="margin-small" href={'/checkin/' + this.state.locationCode} external></Button>
            </Block>
        )
    }
}
