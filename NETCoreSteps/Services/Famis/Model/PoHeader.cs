using System;

namespace Famis.Model
{
    public class PoHeader
    {
        public int? Id { get; set; }

        public string Number { get; set; }

        public string BlanketId { get; set; }

        public string ReleaseNumber { get; set; }

        public int RequestorId { get; set; }

        public string RequestorExternalId { get; set; }

        public string RequestorName { get; set; }

        public string RequestorEmail { get; set; }

        public string RequestorPhone { get; set; }

        public int TypeId { get; set; }

        public int StatusId { get; set; }

        public int? AttentionToId { get; set; }

        public string AttentionToExternalId { get; set; }

        public string BillToAddress { get; set; }

        public int? DefaultRequestId { get; set; }

        public string DefaultRequestExternalId { get; set; }

        public string Description { get; set; }

        public int? VendorId { get; set; }

        public string VendorExternalId { get; set; }

        public string VendorName { get; set; }

        public string VendorAddress { get; set; }

        public DateTimeOffset Date { get; set; }

        public int? WarehouseId { get; set; }

        public string WarehouseExternalId { get; set; }

        public int? PropertyId { get; set; }

        public string PropertyExternalId { get; set; }

        public string ShipToAddress { get; set; }

        public decimal SpendLimit { get; set; }

        public int? ContractId { get; set; }
        
        public bool ShouldSerializeId(){
            return false;
        }
        
        public bool ShouldSerializeStatusId() {
            return false;
        }
        public bool ShouldSerializeVendorName() {
            return false;
        }
    }
}