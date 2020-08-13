// @flow
import { Component } from 'react';
import cookie from 'react-cookies'

type Props = {
    callbackSession: Function
}
type State = {
    session: Object
}

export default class SessionComponent extends Component<Props, State> {
    $f7: any
    $f7router: any
    $f7route: any

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
        };
    }

    componentDidMount() {
        fetch('/auth/status', {
            method: 'GET'
        })
            .then((response) => response.json())
            .then(async (json) => {
                this.setState({
                    session: {
                        useronline: json.useronline,
                        username: json.username                    }
                });
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

    sessionCallback = (sessionData: Object) => {
        this.setState({
            session: {
                useronline: sessionData.useronline,
                username: sessionData.username,
                avatarurl: sessionData.avatarurl,
                connected: true
            }
        });
        if (!sessionData.useronline) {
            // this.$f7router.navigate('/login?callbackUrl=' + this.$f7route.path, { animate: false });
        } else {
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
