import aspida from "@aspida/axios";
import api from "../api/$api";
import axios from "axios";

/** @package */
export const client = api(aspida(axios, { baseURL: process.env.NEXT_PUBLIC_GO_BASE_URL }));

