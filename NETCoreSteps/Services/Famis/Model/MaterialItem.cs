using System;
using System.Collections.Generic;

namespace Famis.Model
{
    public class MaterialItem
    {
        public int? Id { get; set; }
        public string PartNumber { get; set; }
        public string Description { get; set; }
        public string Fin { get; set; }
        public object MaterialClassId { get; set; }
        public float UnitCost { get; set; }
        public string ContraAccountNumber { get; set; }
        public string GlAccountNumber { get; set; }
        public string Manufacturer { get; set; }
        public string Model { get; set; }
        public bool LotControlledFlag { get; set; }
        public int LotShelfLifeDays { get; set; }
        public object ConditionId { get; set; }
        public object Size { get; set; }
        public bool IgnoreMarkupFlag { get; set; }
        public string Barcode { get; set; }
        public string Comments { get; set; }
        public int UnitOfMeasureId { get; set; }
        public object ExternalId { get; set; }
        public bool ActiveFlag { get; set; }
        public bool StockFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public object MaterialCode { get; set; }
        public object MaterialsListId { get; set; }
        public object QuantityOnHand { get; set; }
        public object ReorderQuantity { get; set; }
        public object ReorderPoint { get; set; }
        public object LastCountDate { get; set; }
        public object TabOrder { get; set; }
        public bool InventoryFlag { get; set; }
        public int UpdateById { get; set; }
        public string UpdatedByExternalId { get; set; }
        public object UPCCode { get; set; }
        public object OEMPartNumber { get; set; }
        public object COAAccountId { get; set; }
        public bool LeedCertifiedFlag { get; set; }
        public bool DirectPartFlag { get; set; }
        public bool MSDSFlag { get; set; }
        public object ExternalSystemId { get; set; }
        public object ABCClass { get; set; }
        public object CommodityCodeId { get; set; }
        public object PropertyId { get; set; }
        public object PropertyExternalId { get; set; }
        public object CompanyId { get; set; }
        public object CompanyExternalId { get; set; }
        public List<Suppliermaterial> SupplierMaterials { get; set; }
    }

    public class Suppliermaterial
    {
        public int Id { get; set; }
        public int ItemId { get; set; }
        public int CompanyId { get; set; }
        public float UnitCost { get; set; }
        public string PartNumber { get; set; }
        public bool PrimarySupplierFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public int UpdatedbyId { get; set; }
        public bool Active { get; set; }
        public int UomId { get; set; }
        public int UomConversion { get; set; }
        public string Comments { get; set; }
        public object MinimumReorderQuantity { get; set; }
        public object LeadTime { get; set; }
        public CompanyDetails CompanyDetails { get; set; }
    }

    public class CompanyDetails
    {
        public string odatatype { get; set; }
        public int Id { get; set; }
        public string Name { get; set; }
        public string Addr1 { get; set; }
        public string City { get; set; }
        public string Zip { get; set; }
        public int StateId { get; set; }
        public string State { get; set; }
        public object CountryId { get; set; }
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
        public object CategoryId { get; set; }
        public object SecondaryCategoryId { get; set; }
        public string SicCode { get; set; }
        public string InternalVendorCode { get; set; }
        public string TaxpayerId { get; set; }
        public object ContractTypeId { get; set; }
        public string ContractComments { get; set; }
        public string MobilePhone { get; set; }
        public string InternalVendorCode2 { get; set; }
        public string RiskRating { get; set; }
        public object TypeOfAccessId { get; set; }
        public object PaymentTermId { get; set; }
        public object ShippingMethodId { get; set; }
        public object FreeOnBoardId { get; set; }
        public string RoutingNumber { get; set; }
        public string Addr3 { get; set; }
        public string RemAddr1 { get; set; }
        public string RemAddr2 { get; set; }
        public string RemAddr3 { get; set; }
        public string RemCity { get; set; }
        public string RemZip { get; set; }
        public object RemStateId { get; set; }
        public string Description { get; set; }
        public bool VisitAutoCreateFlag { get; set; }
        public bool DebtorFlag { get; set; }
        public bool LandOwnerFlag { get; set; }
        public bool MeterSiteFlag { get; set; }
        public bool ExtMasterCompanyFlag { get; set; }
    }
}       