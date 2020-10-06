// @flow
import cookie from 'react-cookies'
import type { Session } from '../../js/types';


const state = {
    useronline: false,
    username: '',
    avatarurl: '',
    connected: false
};

const updateSession = (): void => {
    fetch('/auth/status', {
        method: 'GET'
    })
        .then((response) => response.json())
        .then(async (json) => {
            state.useronline = json.useronline;
            state.username = json.username;
            state.avatarurl = json.avatarurl;
            state.connected = true;
        });
    if (state.useronline) {
        cookie.save('lastUser', state.username, { path: '/' });
    }
};

const isUserOnline = (): boolean => {
    return state.useronline;
}

const isSessionConnected = (): boolean => {
    return state.connected;
}

const getSession = () :Session => {
    return state;
}


export { updateSession, isUserOnline, isSessionConnected, getSession };
