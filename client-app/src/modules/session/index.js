// @flow
import { useState, useEffect } from 'react';
import cookie from 'react-cookies'
import type { Session } from '../../js/types';


const initialState: Session = {
    useronline: false,
    username: '',
    avatarurl: '',
    connected: false,
};

const useSession = (): Session => {
    const [session: Session, setSession] = useState(initialState);

    useEffect(() => {
        if (session.connected) {
            return;
        }
        // Session not yet connected: try to get the auth-status from server and build session-object
        fetch('/auth/status', {
            method: 'GET'
        })
            .then((response) => response.json())
            .then(async (json) => {
                setSession({
                    useronline: json.useronline,
                    username: json.username,
                    avatarurl: json.avatarurl,
                    connected: true
                })
                
            });
        if (session.useronline) {
            cookie.save('lastUser', session.username, { path: '/' });
        }
    });

    // const isUserOnline = (): boolean => {
    //     return session.useronline;
    // }

    // const isSessionConnected = (): boolean => {
    //     return session.connected;
    // }

    return session;
};

export { useSession };
