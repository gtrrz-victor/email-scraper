import { Bookings } from "../models/Bookings";

export const fetchBookings = async (): Promise<Bookings[]> => {
    const res = await fetch("https://externalbookings.bqlever.com/bookings", { mode: "cors" });
    return res.json();
};