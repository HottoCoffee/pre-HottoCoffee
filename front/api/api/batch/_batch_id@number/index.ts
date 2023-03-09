/* eslint-disable */
import type * as Types from '../../../@types'

export type Methods = {
  delete: {
    status: 200
    /** successfully deleted */
    resBody: Types.External2_Success
  }

  put: {
    status: 200
    /** successfully updated. */
    resBody: Types.External0_Batch
    reqBody: Types.External0_Batch
  }

  get: {
    status: 200
    /** return specific batch information */
    resBody: Types.External2_Success
  }
}
