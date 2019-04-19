using System;
using Newtonsoft.Json;

namespace Famis.Model
{
    [JsonObject(ItemNullValueHandling = NullValueHandling.Ignore)]
    public class User
    {
        public int? Id { get; set; }
        public int? CompanyId { get; set; }
        public string FirstName { get; set; }
        public string LastName { get; set; }
        public string Title { get; set; }
        public int? TypeId { get; set; }
        public string Addr1 { get; set; }
        public string Addr2 { get; set; }
        public string City { get; set; }
        public string State { get; set; }
        public string Zip { get; set; }
        public string Country { get; set; }
        public string BusPhone { get; set; }
        public string MobPhone { get; set; }
        public string Fax { get; set; }
        public string Email { get; set; }
        public string MobEmail { get; set; }
        public string HomePhone { get; set; }
        public string AsstPhone { get; set; }
        public string UserName { get; set; }
        public string Password { get; set; }
        public bool? ActiveFlag { get; set; }
        public DateTime? UpdateDate { get; set; }
        public int? RequestHistoryDays { get; set; }
        public decimal? RegHourlyRate { get; set; }
        public string ExternalId { get; set; }
        public int? UpdatedById { get; set; }
        public bool? WoAuthFlag { get; set; }
        public string WoAuthComments { get; set; }
        public bool? VisitorEmailFlag { get; set; }
        public bool? EmailWoConfirmationFlag { get; set; }
        public bool? TimeCardFlag { get; set; }
        public int? DefOriginationCodeId { get; set; }
        public int? RequestFutureDays { get; set; }
        public string CubeNumber { get; set; }
        public bool? RestrictedFullUserFlag { get; set; }
        public bool? WorkStatusFlag { get; set; }
        public string DefaultPage { get; set; }
        public string DepartmentDescription { get; set; }
        public int? DepartmentId { get; set; }
        public string PositionDescription { get; set; }
        public int? PositionId { get; set; }
        public int? PositionStandardId { get; set; }
        public int? ExternalSystemId { get; set; }
        public int? ProfileId { get; set; }
        public int? LanguageId { get; set; }
        public bool? UseRateScheduleFlag { get; set; }
        public int? TimeCardFormatId { get; set; }
        public string PayrollExternalId { get; set; }
        public int? RequestsPerPage { get; set; }
        public bool? SelfRegistrationProfileFlag { get; set; }
        public bool? MarkupFlag { get; set; }
        public int? MobileRequestFutureDays { get; set; }
        public int? MobileRequestHistoryDays { get; set; }
        public int? MobileRequestsPerPage { get; set; }
        public DateTime? CoiExpirationDate { get; set; }
        public int? CountryId { get; set; }
        public int? StateId { get; set; }
        public int? AccountId { get; set; }
        public int? DateFormatId { get; set; }
        public int? MobileDateFormatId { get; set; }
        public int? WoApprovalLevelId { get; set; }
        public int? PoApprovalLevelId { get; set; }
        public int? ProjectApprovalLevelId { get; set; }
        public int? PrApprovalId { get; set; }
        public int? PrimaryTimeCardApproverId { get; set; }
        public bool? PrBuyerFlag { get; set; }
        public bool? ProfileFlag { get; set; }
        public bool? PasswordNeverExpiresFlag { get; set; }
        public bool? SsoRequiredFlag { get; set; }
        public int? AlsIncorrectLoginCount { get; set; }
        public bool? AlsForcePasswordChangeFlag { get; set; }
        public bool? AlsNeverInactivateFlag { get; set; }
        public DateTime? AlsLoginExpirationDate { get; set; }
        public int? AlsLoginStatus { get; set; }
        public string CompanyExternalId { get; set; }
        public string ProfileExternalId { get; set; }
        public string UpdatedByExternalId { get; set; }
        public string Name { get; set; }
    }

    public class Companytype
    {
        public int Id { get; set; }
        public string Desc { get; set; }
        public bool? OccupantFlag { get; set; }
    }

    public class Currency
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Abbreviation { get; set; }
        public int Code { get; set; }
        public bool? ActiveFlag { get; set; }
        public bool? InstalledFlag { get; set; }
        public int CurrencyInstallId { get; set; }
        public string Sign { get; set; }
    }

    public class Companycategory
    {
        public int Id { get; set; }
        public string Desc { get; set; }
        public DateTime UpdateDate { get; set; }
        public int UpdatedById { get; set; }
        public bool? ActiveFlag { get; set; }
    }

    public class Secondarycategory
    {
        public int Id { get; set; }
        public string Desc { get; set; }
        public DateTime UpdateDate { get; set; }
        public int UpdatedById { get; set; }
        public bool? ActiveFlag { get; set; }
    }

    public class Paymentterm
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public int TabOrder { get; set; }
        public bool? Active { get; set; }
        public string UpdatedByName { get; set; }
        public DateTime UpdateDate { get; set; }
    }

    public class Shippingmethod
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public int TabOrder { get; set; }
        public bool? ActiveFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedByName { get; set; }
    }

    public class Typeofaccess
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public bool? DefaultSupervisedFlag { get; set; }
        public int TabOrder { get; set; }
        public bool? Active { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdateByName { get; set; }
    }

    public class Companyfreeonboard
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public int TabOrder { get; set; }
        public bool? Active { get; set; }
        public string UpdatedByName { get; set; }
        public DateTime UpdateDate { get; set; }
    }

    public class Profile
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string ExternalId { get; set; }
        public int UserTypeId { get; set; }
    }

    public class Department
    {
        public string Id { get; set; }
        public Number Number { get; set; }
        public Description Description { get; set; }
        public int AuthorizingManagerId { get; set; }
        public string AuthorizingManagerExternalId { get; set; }
        public int CompanyId { get; set; }
        public string CompanyExternalId { get; set; }
        public string OrganizationUnitCode { get; set; }
        public bool? Active { get; set; }
        public int TabOrder { get; set; }
        public int ParentDepartmentId { get; set; }
        public DateTime UpdateDate { get; set; }
        public int UpdatedById { get; set; }
        public string UpdatedByExternalId { get; set; }
    }

    public class Number
    {
    }

    public class Description
    {
    }

    public class Usertype
    {
        public int Id { get; set; }
        public string Description { get; set; }
        public bool? GuestFlag { get; set; }
        public bool? UnnamedGuestFlag { get; set; }
    }

    public class Projectapprovallevel
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public int MaximumAmount { get; set; }
        public int TabOrder { get; set; }
        public bool? PurchaseOrderFlag { get; set; }
        public bool? WorkOrderFlag { get; set; }
        public bool? ProjectFlag { get; set; }
        public bool? PurchaseReqFlag { get; set; }
        public bool? ActiveFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedByName { get; set; }
    }

    public class Poapprovallevel
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public int MaximumAmount { get; set; }
        public int TabOrder { get; set; }
        public bool? PurchaseOrderFlag { get; set; }
        public bool? WorkOrderFlag { get; set; }
        public bool? ProjectFlag { get; set; }
        public bool? PurchaseReqFlag { get; set; }
        public bool? ActiveFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedByName { get; set; }
    }

    public class Prapprovallevel
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public int MaximumAmount { get; set; }
        public int TabOrder { get; set; }
        public bool? PurchaseOrderFlag { get; set; }
        public bool? WorkOrderFlag { get; set; }
        public bool? ProjectFlag { get; set; }
        public bool? PurchaseReqFlag { get; set; }
        public bool? ActiveFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedByName { get; set; }
    }

    public class Woapprovallevel
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public int MaximumAmount { get; set; }
        public int TabOrder { get; set; }
        public bool? PurchaseOrderFlag { get; set; }
        public bool? WorkOrderFlag { get; set; }
        public bool? ProjectFlag { get; set; }
        public bool? PurchaseReqFlag { get; set; }
        public bool? ActiveFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedByName { get; set; }
    }

    public class Laborentryids
    {
    }

    public class Manager
    {
        public int Level { get; set; }
        public int Id { get; set; }
        public string FirstName { get; set; }
        public string LastName { get; set; }
    }
}