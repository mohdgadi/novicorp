# novicorp
This is a simple application where departmental stores can provide dynamic discounts on various items. Items JSON file indicates the preference and the combination of discounts offered.

# Example
Prices:
 Code         | Name                |  Price
-------------------------------------------------
VOUCHER      | NoviCap Voucher      |   5.00€
TSHIRT       | NoviCap T-Shirt      |  20.00€
MUG          | NoviCap Coffee Mug   |   7.50€

Promotion 1: On buying 2 vouchers one is free
Promotion 2: On buying 3 or more Tshirts individual price will be $19.

Examples:

    Items: VOUCHER, TSHIRT, MUG
    Total: 32.50€

    Items: VOUCHER, TSHIRT, VOUCHER
    Total: 25.00€

    Items: TSHIRT, TSHIRT, TSHIRT, VOUCHER, TSHIRT
    Total: 81.00€

    Items: VOUCHER, TSHIRT, VOUCHER, VOUCHER, MUG, TSHIRT, TSHIRT
    Total: 74.50€

Some of the questions answered for design decisions:
https://docs.google.com/document/d/1PjCZHgc5qlK0WAiLrOC9486pAzv5H_nDKy9NQT1E0PU/edit?usp=sharing