import axios from 'axios';
import { ICountry, IState } from '../types';

export class MapApi {
    private static readonly _mapApi = '/api/map/';
    
    public static async getCountries(): Promise<[ICountry]> {
        const url = `${this._mapApi}country`;
        const response = await axios.get(url);
        return response.data
    }
    
    public static async getAllStates(): Promise<[IState]> {
        const url = `${this._mapApi}state`;
        const response = await axios.get(url);
        return response.data
    }
    
    public static async getStates(country: string): Promise<[IState]> {
        const url = `${this._mapApi}state/${country}`;
        const response = await axios.get(url);
        return response.data
    }
}