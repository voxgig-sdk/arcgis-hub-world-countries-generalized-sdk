
import { Context } from './Context'


class ArcgisHubWorldCountriesGeneralizedError extends Error {

  isArcgisHubWorldCountriesGeneralizedError = true

  sdk = 'ArcgisHubWorldCountriesGeneralized'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  ArcgisHubWorldCountriesGeneralizedError
}

