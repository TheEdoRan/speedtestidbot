import axios from "axios";
import memoize from "memoizee";

const searchByName = (query) =>
  axios.get(`https://www.speedtest.net/api/js/servers?search=${query}`);

// Memoize for 12 hours.
const memoOpts = { promise: true, maxAge: 43200 * 1000 };

export const memoSearchByName = memoize(searchByName, memoOpts);
