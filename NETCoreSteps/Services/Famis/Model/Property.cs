using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model
{
    public class Property
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Addr1 { get; set; }
        public string City { get; set; }
        public string Zip { get; set; }
        public int StateId { get; set; }
        public string State { get; set; }
        public int CountryId { get; set; }
        public string Country { get; set; }
        public int TypeId { get; set; }
        public string Phone { get; set; }
        public bool ActiveFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public string ExternalId { get; set; }
        public bool TimeCardFlag { get; set; }
        public bool VendorFlag { get; set; }
        public bool MinorityFlag { get; set; }
        public bool WomanOwnedFlag { get; set; }
        public bool PreferredVendorFlag { get; set; }
        public bool SupplierFlag { get; set; }
        public bool SubcontractorAuthFlag { get; set; }
        public bool W9OnFileFlag { get; set; }
        public int CurrencyInstallId { get; set; }
        public string Addr2 { get; set; }
        public string Fax { get; set; }
        public string Website { get; set; }
        public string EmergencyPhone { get; set; }
        public string Email { get; set; }
        public string PagerNumber { get; set; }
        public string PrimaryContactName { get; set; }
        public int? CategoryId { get; set; }
        public int? SecondaryCategoryId { get; set; }
        public string SicCode { get; set; }
        public string InternalVendorCode { get; set; }
        public string TaxpayerId { get; set; }
        public int? ContractTypeId { get; set; }
        public string ContractComments { get; set; }
        public string MobilePhone { get; set; }
        public string InternalVendorCode2 { get; set; }
        public string RiskRating { get; set; }
        public int? TypeOfAccessId { get; set; }
        public int? PaymentTermId { get; set; }
        public int? ShippingMethodId { get; set; }
        public int? FreeOnBoardId { get; set; }
        public string RoutingNumber { get; set; }
        public string Addr3 { get; set; }
        public string RemAddr1 { get; set; }
        public string RemAddr2 { get; set; }
        public string RemAddr3 { get; set; }
        public string RemCity { get; set; }
        public string RemZip { get; set; }
        public int? RemStateId { get; set; }
        public string Description { get; set; }
        public bool VisitAutoCreateFlag { get; set; }
        public bool DebtorFlag { get; set; }
        public bool LandOwnerFlag { get; set; }
        public bool MeterSiteFlag { get; set; }
        public bool ExtMasterCompanyFlag { get; set; }
        public Account Account { get; set; }
        public List<Udf> Udfs { get; set; }
    }

    public class Account
    {
        public int Id { get; set; }
        public string Description { get; set; }
        public int UpdatedById { get; set; }
        public string UpdatedByExternalId { get; set; }
        public DateTime UpdateDate { get; set; }
        public bool ActiveFlag { get; set; }
        public int MaskId { get; set; }
    }
}
