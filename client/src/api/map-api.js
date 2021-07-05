import axios from "axios";

const _mapApi = "/api/map/";

async function getCountries() {
  const url = `${_mapApi}country`;
  const response = await axios.get(url);
  return response.data;
}

async function getAllStates() {
  const url = `${_mapApi}state`;
  const response = await axios.get(url);
  return response.data;
}

async function getStates(country) {
  const url = `${_mapApi}state/${country}`;
  const response = await axios.get(url);
  return response.data;
}

export { getCountries, getAllStates, getStates };
