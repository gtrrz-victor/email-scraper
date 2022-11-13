import * as React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { useQuery } from "react-query";
import { fetchBookings } from '../dataProvider/fetchBookings';
import { Button, CircularProgress, Stack } from '@mui/material';
import { Bookings } from '../models/Bookings';

export default function BasicTable() {
    const { data, status, refetch } = useQuery<Bookings[]>("bookings", fetchBookings);

    if (status === "loading") {
        return (
            <CircularProgress />
        )
    }
    return (
        <Stack spacing={2}>
            <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Tour</TableCell>
                            <TableCell align="right">Date</TableCell>
                            <TableCell align="right">Name</TableCell>
                            <TableCell align="right">Number of people</TableCell>
                            <TableCell align="right">Phone number</TableCell>
                            <TableCell align="right">Email</TableCell>
                            <TableCell align="right">Booking number</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {data && data.map((row) => (
                            <TableRow
                                key={row.name}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell component="th" scope="row">
                                    {row.tour}
                                </TableCell>
                                <TableCell align="right">{row.date}</TableCell>
                                <TableCell align="right">{row.name}</TableCell>
                                <TableCell align="right">{row.numPeople}</TableCell>
                                <TableCell align="right">{row.phoneNumber}</TableCell>
                                <TableCell align="right">{row.contactEmail}</TableCell>
                                <TableCell align="right">{row.bookingNumber}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
            <Button variant="contained" onClick={() => refetch()}>Reload Content</Button>
        </Stack>
    );
}