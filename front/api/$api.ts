import type { AspidaClient, BasicHeaders } from 'aspida'
import { dataToURLString } from 'aspida'
import type { Methods as Methods0 } from './api/batch'
import type { Methods as Methods1 } from './api/batch/_id@number/history'
import type { Methods as Methods2 } from './api/batch/_id@number/history/_historyId@number'
import type { Methods as Methods3 } from './api/calendar'

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? 'http://localhost:3000' : baseURL).replace(/\/$/, '')
  const PATH0 = '/api/batch'
  const PATH1 = '/history'
  const PATH2 = '/api/calendar'
  const GET = 'GET'
  const POST = 'POST'

  return {
    api: {
      batch: {
        _id: (val2: number) => {
          const prefix2 = `${PATH0}/${val2}`

          return {
            history: {
              _historyId: (val4: number) => {
                const prefix4 = `${prefix2}${PATH1}/${val4}`

                return {
                  /**
                   * @returns Success to get specific history
                   */
                  get: (option?: { config?: T | undefined } | undefined) =>
                    fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, prefix4, GET, option).json(),
                  /**
                   * @returns Success to get specific history
                   */
                  $get: (option?: { config?: T | undefined } | undefined) =>
                    fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, prefix4, GET, option).json().then(r => r.body),
                  $path: () => `${prefix}${prefix4}`
                }
              },
              /**
               * @returns Success to get history
               */
              get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods1['get']['resBody'], BasicHeaders, Methods1['get']['status']>(prefix, `${prefix2}${PATH1}`, GET, option).json(),
              /**
               * @returns Success to get history
               */
              $get: (option?: { config?: T | undefined } | undefined) =>
                fetch<Methods1['get']['resBody'], BasicHeaders, Methods1['get']['status']>(prefix, `${prefix2}${PATH1}`, GET, option).json().then(r => r.body),
              /**
               * @returns Success to receive
               */
              post: (option: { body: Methods1['post']['reqBody'], config?: T | undefined }) =>
                fetch<Methods1['post']['resBody'], BasicHeaders, Methods1['post']['status']>(prefix, `${prefix2}${PATH1}`, POST, option).json(),
              /**
               * @returns Success to receive
               */
              $post: (option: { body: Methods1['post']['reqBody'], config?: T | undefined }) =>
                fetch<Methods1['post']['resBody'], BasicHeaders, Methods1['post']['status']>(prefix, `${prefix2}${PATH1}`, POST, option).json().then(r => r.body),
              $path: () => `${prefix}${prefix2}${PATH1}`
            }
          }
        },
        /**
         * Get batch list. Response can be filtered by keyword.
         * @returns Success filter and can get batch list
         */
        get: (option?: { query?: Methods0['get']['query'] | undefined, config?: T | undefined } | undefined) =>
          fetch<Methods0['get']['resBody'], BasicHeaders, Methods0['get']['status']>(prefix, PATH0, GET, option).json(),
        /**
         * Get batch list. Response can be filtered by keyword.
         * @returns Success filter and can get batch list
         */
        $get: (option?: { query?: Methods0['get']['query'] | undefined, config?: T | undefined } | undefined) =>
          fetch<Methods0['get']['resBody'], BasicHeaders, Methods0['get']['status']>(prefix, PATH0, GET, option).json().then(r => r.body),
        /**
         * Register new batch
         * @returns Success to create
         */
        post: (option: { body: Methods0['post']['reqBody'], config?: T | undefined }) =>
          fetch<Methods0['post']['resBody'], BasicHeaders, Methods0['post']['status']>(prefix, PATH0, POST, option).json(),
        /**
         * Register new batch
         * @returns Success to create
         */
        $post: (option: { body: Methods0['post']['reqBody'], config?: T | undefined }) =>
          fetch<Methods0['post']['resBody'], BasicHeaders, Methods0['post']['status']>(prefix, PATH0, POST, option).json().then(r => r.body),
        $path: (option?: { method?: 'get' | undefined; query: Methods0['get']['query'] } | undefined) =>
          `${prefix}${PATH0}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`
      },
      calendar: {
        /**
         * @returns Return batch history list
         */
        get: (option?: { query?: Methods3['get']['query'] | undefined, config?: T | undefined } | undefined) =>
          fetch<Methods3['get']['resBody'], BasicHeaders, Methods3['get']['status']>(prefix, PATH2, GET, option).json(),
        /**
         * @returns Return batch history list
         */
        $get: (option?: { query?: Methods3['get']['query'] | undefined, config?: T | undefined } | undefined) =>
          fetch<Methods3['get']['resBody'], BasicHeaders, Methods3['get']['status']>(prefix, PATH2, GET, option).json().then(r => r.body),
        $path: (option?: { method?: 'get' | undefined; query: Methods3['get']['query'] } | undefined) =>
          `${prefix}${PATH2}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`
      }
    }
  }
}

export type ApiInstance = ReturnType<typeof api>
export default api