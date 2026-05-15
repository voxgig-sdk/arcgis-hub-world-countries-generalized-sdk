
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { ArcgisHubWorldCountriesGeneralizedSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await ArcgisHubWorldCountriesGeneralizedSDK.test()
    equal(null !== testsdk, true)
  })

})
