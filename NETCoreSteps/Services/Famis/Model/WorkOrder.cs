using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using Newtonsoft.Json;

namespace Famis.Model
{
    public class WorkOrder
    {
        [JsonProperty(PropertyName = "id")]
        public int? Id { get; set; }
        public string RequestTypeId { get; set; }
        public string RequestSubTypeId { get; set; }
        public string RequestPriorityId { get; set; }
        public string StatusId { get; set; }
        public string LastStatusEnum { get; set; }
        public string StatementOfWork { get; set; }
        public string ExternalId { get; set; }
        public string ExternalId2 { get; set; }
        public string AssignedDate { get; set; }
        public string ClosedDate { get; set; }
        public string ClosedById { get; set; }
        public string CreateDate { get; set; }
        public string CreatedById { get; set; }
        public string CompleteByDate { get; set; }
        public string UpdateDate { get; set; }
        public string DateScheduled { get; set; }
        public string TimeScheduled { get; set; }
        public string ActualCompletionDate { get; set; }
        public string SlaEstimatedResponseDate { get; set; }
        public string SlaActualResponseDate { get; set; }
        public string SlaEstimatedCompletionDate { get; set; }
        public string SlaActualCompletionDate { get; set; }
        public bool SlaOverdueAlertFlag { get; set; }
        public string PropertyHoursInDay { get; set; }
        public string ActualCompletionTime { get; set; }
        public string ActualResponseTime { get; set; }
        public bool ExcludeSlaReportingFlag { get; set; }
        public bool ResponseExcludedFlag { get; set; }
        public bool CompletionExcludedFlag { get; set; }
        public string ReassignReasonId { get; set; }
        public string CompletionReasonId { get; set; }
        public string RequestedCompletionDate { get; set; }
        public string SpaceId { get; set; }
        public string SubSpaceId { get; set; }
        public string RoomCube { get; set; }
        public string NotifyAssignedToMethod { get; set; }
        public bool CloseNotificationFlag { get; set; }
        public bool NotifyAssignedToFlag { get; set; }
        public bool ReNotifyAssignedToFlag { get; set; }
        public string NotifyFollowUpAlert { get; set; }
        public bool InitialNotificationFlag { get; set; }
        public bool FollowUpNotificationFlag { get; set; }
        public string NotifyDate { get; set; }
        public string EmailCC { get; set; }
        public bool NotifyRequestorAutoSentFlag { get; set; }
        public bool NotifySlaReminderSentFlag { get; set; }
        public string AssignedToId { get; set; }
        public int? CrewId { get; set; }
        public string RequestorLastName { get; set; }
        public string RequestorFirstName { get; set; }
        public string RequestorName { get; set; }
        public string RequestorPhone { get; set; }
        public string RequestorEmail { get; set; }
        public string RequestorFax { get; set; }
        public string RequestorCompanyName { get; set; }
        public string RequestorId { get; set; }
        public string CreatedByFirstName { get; set; }
        public string CreatedByLastName { get; set; }
        public string CreatedByPhone { get; set; }
        public string CreatedByEmail { get; set; }
        public bool BillableFlag { get; set; }
        public bool BillingStatusFlag { get; set; }
        public string TotalLaborCost { get; set; }
        public string TotalMaterialCost { get; set; }
        public string TotalOtherCost { get; set; }
        public string TotalMarkup { get; set; }
        public string BillCodeId { get; set; }
        public string CostCodeId { get; set; }
        public string NotToExceedAmount { get; set; }
        public string NotToExceedComment { get; set; }
        public string EstimatedLaborHours { get; set; }
        public string ExternalCostCenter { get; set; }
        public string ExternalCostCenterId { get; set; }
        public string ExternalCostCenterDescription { get; set; }
        public string EstimatedTotalAmount { get; set; }
        public string TotalLaborHours { get; set; }
        public string DefaultAccountId { get; set; }
        public string DefaultCoaAccountId { get; set; }
        public string DefaultCoaCreditAccountId { get; set; }
        public string TotalTax { get; set; }
        public string ServiceProviderRefNumber { get; set; }
        public string ServiceProviderToInvoiceFlag { get; set; }
        public string ChargeTypeId { get; set; }
        public string AssetNumber { get; set; }
        public string AssetId { get; set; }
        public string FailureCodeId { get; set; }
        public string RequireAssetFlag { get; set; }
        public string InspectionId { get; set; }
        public string InspectionDetailId { get; set; }
        public string RecurrenceId { get; set; }
        public string ProcedureId { get; set; }
        public string ProcedureNameHistory { get; set; }
        public string AttachedFileName { get; set; }
        public string ProcedureBodyHistory { get; set; }
        public string ProvisionId { get; set; }
        public string ProvisionDetailId { get; set; }
        public string IncidentId { get; set; }
        public string ReservationId { get; set; }
        public string ParentWOId { get; set; }
        public bool TopLevelFlag { get; set; }
        public string EscalateToId { get; set; }
        public bool ServiceEscalationFlag { get; set; }
        public string ApprovedByName { get; set; }
        public string ApprovedDate { get; set; }
        public bool RequestApprovalflag { get; set; }
        public bool AuthentryFlag { get; set; }
        public string AuthentryRemarks { get; set; }
        public string CustomerPoNumber { get; set; }
        public string InvoiceDate { get; set; }
        public string ProjectId { get; set; }
        public string ExternalCompanyDescription { get; set; }
        public string CompanyDocumentId { get; set; }
        public string VendorCompanyId { get; set; }
        public bool VendorUserFlag { get; set; }
        public string VendorInvoiceCompany { get; set; }
        public string VendorInvoiceAddress { get; set; }
        public string VendorInvoicePhone { get; set; }
        public string VendorInvoiceFax { get; set; }
        public string DepartmentId { get; set; }
        public bool KbaseFlag { get; set; }
        public string ContractNumber { get; set; }
        public string DispatchDate { get; set; }
        public string ResolutionCodeId { get; set; }
        public string OriginationCodeId { get; set; }
        public bool ExportFlag { get; set; }
        public string ExportDate { get; set; }
        public string CloseByDateChangeCount { get; set; }
        public bool CorrectiveRequestFlag { get; set; }
        public string ExternalSystemId { get; set; }
        public string DispatchExternalAckDate { get; set; }
        public string CloseExternalAckDate { get; set; }
        public string InvoiceNumber { get; set; }
        public string AltPhone { get; set; }
        public string RoomTypeId { get; set; }
        public string Signature { get; set; }
        public string SignatureText { get; set; }
        public string SignatureDate { get; set; }
        public string BudgetYear { get; set; }
        public bool BudgetFlag { get; set; }
        public string RunTimeValue { get; set; }
        public string RunTimeTypeId { get; set; }
        public bool BeenOnHoldFlag { get; set; }
        public string AuthorizerName { get; set; }
        public string AuthorizerPhone { get; set; }
        public string EstimatedArrivalDate { get; set; }
        public string ActualArrivalDate { get; set; }
        public string ArBatchId { get; set; }
        public string RequestWaiverComments { get; set; }
        public string RequestWaiverId { get; set; }
        public string ExtInvAmount { get; set; }
        public string ExternalPOLineNumber { get; set; }
        public string TrackingCodeId { get; set; }
        public bool PsArInvoiceFlag { get; set; }
        public string IsOpenRequest { get; set; }
        public string IsClosedRequest { get; set; }
        public bool Pmstringegrated { get; set; }
        public string UserGroupId { get; set; }
        public string WorkTypeId { get; set; }
        public string PropertyId { get; set; }
        public string PropertyExternalId { get; set; }
        public string RequestTypeExternalId { get; set; }
        public string RequestSubTypeExternalId { get; set; }
        public string SpaceExternalId { get; set; }
        public string SubSpaceExternalId { get; set; }
        public string RequestPriorityExternalId { get; set; }
        public string CreatedByExternalId { get; set; }
        public string ClosedByExternalId { get; set; }
        public string AssignedToExternalId { get; set; }
        public string StatusExternalId { get; set; }
        public string InspectionExternalId { get; set; }
        public string AssetExternalId { get; set; }
        public string ProcedureExternalId { get; set; }
        public string EscalateToExternalId { get; set; }
        public string IncidentExternalId { get; set; }
        public string ParentWOExternalId { get; set; }
        public string InspectionDetailExternalId { get; set; }
        public string VendorCompanyExternalId { get; set; }
        public string RequestorExternalId { get; set; }
        public string CrewExternalId { get; set; }
        public string GeneralComments { get; set; }
        public string InternalComments { get; set; }
        public bool NotifyRequestorFlag { get; set; }

        private bool IsActive => Status?.ActiveFlag == true;

        public Status Status { get; set; }

        public virtual Space Space { get; set; }

        public virtual AccountInfo AccountInfo { get; set; }
    }

    public class AccountInfo
    {
        public int Id { get; set; }
        public int WorkOrderId { get; set; }
        public string HoldReason { get; set; }
        public bool HoldFlag { get; set; }
        public int? BillingTypeId { get; set; }

        public virtual List<Detail> Details { get; set; }
    }

    public class Detail
    {
        public int Id { get; set; }
        public int AccountInfoId { get; set; }
        public decimal? Percentage { get; set; }
        public int? ChartOfAccountsId { get; set; }
        public int? IndexId { get; set; }

        public virtual List<Segment> Segments { get; set; }
    }

    public class Segment
    {
        public int? SegmentId { get; set; }
        public int DetailId { get; set; }
        public int? SegmentValueId { get; set; }
        public string SegmentValue { get; set; }
    }

    public class Status
    {
        [JsonProperty("Id")]
        public long Id { get; set; }

        [JsonProperty("ExternalId")]
        public string ExternalId { get; set; }

        [JsonProperty("Name")]
        public string Name { get; set; }

        [JsonProperty("TabOrder")]
        public long TabOrder { get; set; }

        [JsonProperty("ActiveFlag")]
        public bool ActiveFlag { get; set; }

        [JsonProperty("DefaultOpenStatusFlag")]
        public bool DefaultOpenStatusFlag { get; set; }

        [JsonProperty("DefaultCloseStatusFlag")]
        public bool DefaultCloseStatusFlag { get; set; }

        [JsonProperty("DefaultGuestCompleteFlag")]
        public bool DefaultGuestCompleteFlag { get; set; }

        [JsonProperty("DefaultApprovedFlag")]
        public bool DefaultApprovedFlag { get; set; }

        [JsonProperty("DefaultDeclinedFlag")]
        public bool DefaultDeclinedFlag { get; set; }

        [JsonProperty("DefaultWaitingApprovalFlag")]
        public bool DefaultWaitingApprovalFlag { get; set; }

        [JsonProperty("DefaultWorkCompleteFlag")]
        public bool DefaultWorkCompleteFlag { get; set; }

        [JsonProperty("OpenStatusFlag")]
        public bool OpenStatusFlag { get; set; }

        [JsonProperty("ClosingStatusFlag")]
        public bool ClosingStatusFlag { get; set; }

        [JsonProperty("EnteredInErrorFlag")]
        public bool EnteredInErrorFlag { get; set; }

        [JsonProperty("WorkCompleteFlag")]
        public bool WorkCompleteFlag { get; set; }

        [JsonProperty("OnHoldFlag")]
        public bool OnHoldFlag { get; set; }

        [JsonProperty("ResponseFlag")]
        public bool ResponseFlag { get; set; }

        [JsonProperty("InProgressFlag")]
        public bool InProgressFlag { get; set; }

        [JsonProperty("ServiceEscalationFlag")]
        public bool ServiceEscalationFlag { get; set; }

        [JsonProperty("OverrideLaborRequiredFlag")]
        public bool OverrideLaborRequiredFlag { get; set; }

        [JsonProperty("SendCloseNotificationFlag")]
        public bool SendCloseNotificationFlag { get; set; }

        [JsonProperty("LockedFlag")]
        public long LockedFlag { get; set; }

        [JsonProperty("LockedFinancialFlag")]
        public long LockedFinancialFlag { get; set; }

        [JsonProperty("SendWebServicesUpdatesFlag")]
        public bool SendWebServicesUpdatesFlag { get; set; }

        [JsonProperty("CommentIsRequiredOnChange")]
        public bool CommentIsRequiredOnChange { get; set; }

        [JsonProperty("DefaultCcStatusFlag")]
        public bool DefaultCcStatusFlag { get; set; }

        [JsonProperty("DefaultInProgressStatusFlag")]
        public bool DefaultInProgressStatusFlag { get; set; }

        [JsonProperty("DefaultCanceledStatus")]
        public bool DefaultCanceledStatus { get; set; }

        [JsonProperty("StatusEnum")]
        public string StatusEnum { get; set; }

        [JsonProperty("StatusCode")]
        public string StatusCode { get; set; }
    }
}