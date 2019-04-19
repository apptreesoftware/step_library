using System;

namespace Famis.Model
{
    public class PoHeaderRequest
    {
        public string RequestorExternalId { get; set; }
        public string BillToAddress { get; set; }
        public int? DefaultRequestId { get; set; }
        public string Description { get; set; }
        public int VendorId { get; set; }
        public string VendorAddress { get; set; }
        public DateTime Date { get; set; }
        public int? WarehouseId { get; set; }
        public string WarehouseExternalId { get; set; }
        public int PropertyId { get; set; }
        public string ShipToAddress { get; set; }
        public double SpendLimit { get; set; }
        public string TypeName { get; set; }
    }

    public class PoHeaderCreateRequest : PoHeaderRequest
    {
        public string Number { get; set; }
        public int TypeID { get; set; }
    }
}