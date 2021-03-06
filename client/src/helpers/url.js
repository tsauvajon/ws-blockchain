import axios from 'axios'

const baseUrl = 'http://localhost:3333'

const apiPrefix = 'api'
const soccerPrefix = 'soccer'
const blockchainPrefix = 'blockchain'

const getTeamByIdFactory = (base) => (id) => axios.get(`${base}/teams/${id}`)
const getFixturesFactory = (base) => (begin, end) => axios.get(`${base}/fixtures/${begin}/${end}`)
const getFixtureByIdFactory = (base) => (id) => axios.get(`${base}/fixtures/${id}`)
const getOddsByFixtureFactory = (base) => (id) => axios.get(`${base}/odds/fixture/${id}`)
const getHighlightsByFixtureFactory = (base) => (id) => axios.get(`${base}/highlights/fixture/${id}`)
const getStandingsBySeasonIdFactory = (base) => (id) => axios.get(`${base}/standings/season/live/${id}`)

const connectFactory = (base) => (username, password) => axios.get(`${base}/connect/${username}/${password}`)
const getBetsFactory = (base) => () => axios.get(`${base}/bets`)
const postBetFactory = (base) => ({ userID, amount, fixtureID, choice }) => axios.post(`${base}/bets`, { userID, amount, fixtureID, choice })

const soccerFactory = (base) => ({
  getTeamById: getTeamByIdFactory(`${base}/${soccerPrefix}`),
  getFixtures: getFixturesFactory(`${base}/${soccerPrefix}`),
  getOddsByFixture: getOddsByFixtureFactory(`${base}/${soccerPrefix}`),
  getStandingsBySeasonId: getStandingsBySeasonIdFactory(`${base}/${soccerPrefix}`),
  getFixtureById: getFixtureByIdFactory(`${base}/${soccerPrefix}`),
  getHighlightsByFixture: getHighlightsByFixtureFactory(`${base}/${soccerPrefix}`)
})

const blockchainFactory = (base) => ({
  connect: connectFactory(`${base}/${blockchainPrefix}`),
  getBets: getBetsFactory(`${base}/${blockchainPrefix}`),
  postBet: postBetFactory(`${base}/${blockchainPrefix}`)
})

const api = (baseUrl) => ({
  soccer: soccerFactory(`${baseUrl}/${apiPrefix}`),
  blockchain: blockchainFactory(`${baseUrl}/${apiPrefix}`)
})

export default api(baseUrl)
