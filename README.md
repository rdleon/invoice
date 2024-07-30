# Generate an invoice

Given a description, an daily rate and a template generate an
html template for worked days in the current month.

The program calculates the weekdays in the current month,
optionally subtrackts the out of office days and outputs the
total owed for the month.

$ invoice [--out-of-office days] <description> <daily rate>

--out-of-office <N>     Days not worked, and not invoiced
--html                  Use the template to output the result in html
