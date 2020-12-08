// @flow
type Session = {
  useronline: boolean,
  username?: string,
  avatarurl?: string,
  connected: boolean
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
  fetched?: boolean
}

type BusinessInfo = {
  description: string,
  language: string
}

// This is the type definition for the json that returns when /c/userdata is fetched
// uuid is the uuid of the (previous) checkin.
type UserDataResponse = {
  businessAddress?: string,
  businessName?: string,
  businessUUID?: string,
  checkedInAt?: string,
  checkedOutAt?: string,
  userUUID: string,
  usercity: string,
  useremail: string,
  username: string,
  userphone: string,
  userstreet: string,
  usercountry: string,
  uuid:string
}

// Props
type BusinessProps = {
  businessData: BusinessData
}

type CheckinProps = {
  businessData: BusinessData,
  chckrCode: string,
  $f7route: Object,
  $f7router: Object,
  f7route: Object,
  f7router: Object
}

type LogoProps = {
  plain?: boolean,
  direction?: string
}


export type { BusinessData, BusinessInfo, Session, UserDataResponse, BusinessProps, CheckinProps, LogoProps };
