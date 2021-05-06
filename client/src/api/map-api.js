import axios from "axios";

const _mapApi = "/api/map/";

async function getCountries() {
  axios.defaults.baseURL = "localhost:8000/";
  const url = `${_mapApi}country`;
  const response = await axios.get(url);
  return response.data;
}

async function getAllStates() {
  axios.defaults.baseURL = "localhost:8000/";
  const url = `${_mapApi}state`;
  const response = await axios.get(url);
  return response.data;
}

async function getStates(country) {
  axios.defaults.baseURL = "localhost:8000/";
  const url = `${_mapApi}state/${country}`;
  const response = await axios.get(url);
  return response.data;
}

export { getCountries, getAllStates, getStates };
