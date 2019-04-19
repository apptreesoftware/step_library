using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace Famis.Model
{
    public class Crew
    {
        public int Id { get; set; }
        public string ExternalId { get; set; }
        public bool ActiveFlag { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedById { get; set; }
        public string ExternalUpdatedById { get; set; }
        public string Description { get; set; }
        public string Rate { get; set; }
        public string OT { get; set; }
        public string DT { get; set; }
        public object CompanyId { get; set; }
        public object CompanyExternalId { get; set; }
        public object DepartmentId { get; set; }
        public object DepartmentExternalId { get; set; }
        public bool RateScheduleFlag { get; set; }
        public bool IsCrewFlag { get; set; }
        public string CrewHoursDay { get; set; }
    }
}