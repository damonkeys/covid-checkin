// @flow
type Session = {
    useronline?: boolean,
    username?: string,
    avatarurl?: string,
    connected?: boolean,  
  }

  type BusinessData = {
    uuid?: string,
    code?: string,
    name?: string,
    street?: string,
    zip?: string,
    city?: string,
    formattedAddress?: string, 
    businessInfos?: BusinessInfo[],
  }

  type BusinessInfo = {
    description: string,
    language: string,
  }

export type { BusinessData, BusinessInfo, Session };
