// @flow
import { Component } from 'react';
import cookie from 'react-cookies'
import type { Session } from '../../js/types';

type Props = {
    callbackSession: Function
}
type State = {
    session: Session
}

export default class SessionComponent extends Component<Props, State> {
    constructor(props: Object) {
        super(props);
        if (!this.state) {
            this.state = {
                session: {}
            }
        }
        this.state.session = {
            useronline: false,
            username: '',
            avatarurl: '',
            connected: false        // connected is true after the first status request to get infos about onlinestatus. we need connected for avoid "blinking"
            // after loading the first time. without connected you will see the login-buttons for a short time even though your are logged in!
            // Blinking means:
            //
            // 1. Loading site... It loads one part of the side ie. Login Buttons
            //      - BUT your are logged in already!
            // 2. React connects auth to check a valid session
            // 3. Auth answers, alright, user is logged in!
            // 4. React rendered site-parts again and hides login buttons.
        };
    }

    componentDidMount() {
        fetch('/auth/status', {
            method: 'GET'
        })
            .then((response) => response.json())
            .then(async (json) => {
                this.sessionCallback(json);
            });
    }

    clickLogout() {
        fetch('/auth/logout', {
            method: 'GET'
        })
            .then((response) => response.json())
            .then(async (json) => {
                if (json.successful) {
                    window.location.reload();
                }
            });
    }

    sessionCallback = (sessionData: Session) => {
        this.setState({
            session: {
                useronline: sessionData.useronline,
                username: sessionData.username,
                avatarurl: sessionData.avatarurl,
                connected: true
            }
        });
        if (sessionData.useronline) {
            cookie.save('lastUser', sessionData.username, { path: '/' });
        }
    }

    sessionSetter = (data: Object) => {
        this.setState({
            session: {
                useronline: data.useronline || this.state.session.useronline,
                username: data.username || this.state.session.username,
                avatarurl: data.avatarurl || this.state.session.avatarurl,
            }
        })
    }
}
