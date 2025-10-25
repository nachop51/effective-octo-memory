export enum AccountCurrency {
  USD = 'USD',
  EUR = 'EUR',
  GBP = 'GBP',
  MUR = 'MUR',
  UYU = 'UYU',
}

export const accountCurrencies = Object.values(AccountCurrency) as [
  AccountCurrency,
  ...AccountCurrency[],
]

export const DEFAULT_CURRENCY = AccountCurrency.USD
