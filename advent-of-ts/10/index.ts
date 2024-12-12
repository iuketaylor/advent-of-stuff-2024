export enum Gift {
  Coal = 0 << 0,
  Train = 1 << 0,
  Bicycle = 1 << 1,
  SuccessorToTheNintendoSwitch = 1 << 2,   
  TikTokPremium = 1 << 3,
  Vape = 1 << 4,
  Traditional = Train | Bicycle,
  OnTheMove = Coal | Bicycle | TikTokPremium | Vape,
  OnTheCouch = Coal | SuccessorToTheNintendoSwitch | TikTokPremium | Vape
};
