import axios from "axios";

export const searchByName = (query) =>
  axios.get(
    `https://www.speedtest.net/api/js/servers?search=${query}&limit=10`,
  );
